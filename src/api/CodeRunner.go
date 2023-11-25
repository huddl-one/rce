package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.huddl.one/rce/src/docker"
	"go.huddl.one/rce/src/utils"
)

type Payload struct {
	RunID    string `json:"runId"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

type TestResult struct {
	Passed   bool        `json:"passed"`
	Input    interface{} `json:"input"` // Use interface{} to handle various input types
	Expected interface{} `json:"expected"`
	Output   interface{} `json:"output"`
	Error    string      `json:"error"`
}

type DockerOutput struct {
	TestResults []TestResult `json:"testResults"`
}

func cleanAndFormatOutput(rawOutput []byte) (string, error) {
	// Convert raw output to string
	outputString := string(rawOutput)

	// Remove control characters and unneeded parts
	cleanedOutput := removeControlCharacters(outputString)

	// Return the cleaned and formatted JSON string
	return cleanedOutput, nil
}

func removeControlCharacters(output string) string {
	// Regex to match control characters
	re := regexp.MustCompile(`[\x00-\x1F\x7F]|\x1B\[.*?m`)
	// Replace all matches with empty string
	return re.ReplaceAllString(output, "")
}

func ExecuteCode(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var payload Payload

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			panic(err)
		}

		// Print the payload for debugging
		fmt.Println(payload)

		filePath := utils.CreateCodeFile(payload.Code, payload.Language, payload.RunID, "runs")

		runId := payload.RunID

		// Print the file path for debugging
		fmt.Println(filePath)

		// convert the output to json
		// output, err := json.Marshal(map[string]string{
		// 	"output": string(docker.RunContainer(filePath, payload.Language, runId)),
		// 	"runId":  runId,
		// })
		// if err != nil {
		// 	panic(err)
		// }

		containerOutput := docker.RunContainer(filePath, payload.Language, runId)

		var results DockerOutput
		err1 := json.Unmarshal(containerOutput, &results)
		if err1 != nil {
			// Output is not JSON - likely an error message
			handlePlainTextOutput(res, containerOutput)
			return
		}

		jsonResponse, _ := json.Marshal(results)
		formattedOutput, err := cleanAndFormatOutput(jsonResponse)
		if err != nil {
			http.Error(res, "Error formatting output: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// set the content type to json, enable CORS and write the output
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Write([]byte(formattedOutput))
	case http.MethodGet:
		res.Write([]byte("Hello World"))
	}
}

func handlePlainTextOutput(res http.ResponseWriter, output []byte) {
	outputStr := string(output)

	// Check for common syntax error patterns
	if strings.Contains(outputStr, "SyntaxError") || strings.Contains(outputStr, "syntax error") {
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusBadRequest) // Bad Request for syntax errors
		res.Write([]byte("Syntax error detected: " + outputStr))
		return
	}

	// Check for common runtime error patterns
	if strings.Contains(outputStr, "RuntimeError") || strings.Contains(outputStr, "runtime error") {
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusInternalServerError) // Internal Server Error for runtime errors
		res.Write([]byte("Runtime error detected: " + outputStr))
		return
	}

	// If no specific error patterns are identified, respond with a generic error message
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte("Error lmao: " + outputStr))
}

func Serve(PORT string) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001", "https://*.huddl.one", "https://huddl.one"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("RCE is Online"))
	})
	router.HandleFunc("/run", ExecuteCode)
	fmt.Println("Server running on port " + PORT)
	fmt.Println("Visit http://localhost:" + PORT)
	fmt.Println("Press CTRL+C to exit")
	http.ListenAndServe("0.0.0.0:"+PORT, router)
}

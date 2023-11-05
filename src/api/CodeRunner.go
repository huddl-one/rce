package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		output, err := json.Marshal(map[string]string{
			"output": string(docker.RunContainer(filePath, payload.Language, runId)),
			"runId":  runId,
		})
		if err != nil {
			panic(err)
		}

		// set the content type to json, enable CORS and write the output
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Write(output)
	case http.MethodGet:
		res.Write([]byte("Hello World"))
	}
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

package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetLanguageFileExtension(language string) string {
	switch language {
	case "python":
		return ".py"
	case "c":
		return ".c"
	case "cpp":
		return ".cpp"
	case "javascript":
		return ".js"
	case "java":
		return ".java"
	default:
		return ".py"
	}
}

func FormatJavaCode(code string, runId string) string {
	return fmt.Sprintf("package %s; public class Solution { public static void main(String[] args) { %s } }", runId, code)
}

func CreateCodeFile(code string, language string, runId string, dir string) string {

	// create the file name
	folderPath := filepath.Join(dir, runId)

	filename := "solution"

	if language == "java" {
		filename = "Solution"
	}

	// create the file path
	filePath := filepath.Join(folderPath, filename+GetLanguageFileExtension(language))

	if language == "java" {
		code = FormatJavaCode(code, runId)
	}

	// create the runs directory if it doesn't exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0777)
	}

	// create the file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	// close the file when the function returns
	defer file.Close()

	// write the code to the file
	_, err = file.WriteString(code)
	if err != nil {
		panic(err)
	}

	return filePath
}

package docker

import (
	"strings"
)

// getDockerImage returns the docker image required to run code for the given language
func selectDockerImageForLanguage(language string) string {
	switch language {
	case "python":
		return "python:3.8-slim"
	case "c":
		return "gcc:latest"
	case "cpp":
		return "gcc:latest"
	case "javascript":
		return "node:20"
	case "java":
		return "openjdk:11"
	default:
		return "python:3.8-slim"
	}
}

// getRunCommand returns the command to run the code for the given language
func generateRunCommandForLanguage(language string, filePath string) []string {
	// split the file path into the file name including path and the extension
	filePathParts := strings.Split(filePath, ".")

	// return the command to run the code
	switch language {
	case "python":
		return []string{"python", filePath}
	case "c":
		return []string{"bash", "-c", "gcc " + filePath + " -o " + filePathParts[0] + " && " + filePathParts[0]}
	case "cpp":
		return []string{"bash", "-c", "g++ " + filePath + " -o " + filePathParts[0] + " && " + filePathParts[0]}
	case "javascript":
		return []string{"node", filePath}
	case "java":
		return []string{"bash", "-c", "javac " + filePath + " && java " + filePathParts[0]}
	default:
		return []string{"python", filePath}
	}
}

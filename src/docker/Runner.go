package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunContainer(filePath string, language string, runId string) []byte {
	ctx := context.Background()

	// Connect to the Docker daemon
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Negotiate to use compatable Docker API version
	cli.NegotiateAPIVersion(ctx)

	// get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// get the command to run the code
	runCommand := generateRunCommandForLanguage(language, strings.Replace(filePath, "runs/", "", 1))
	fmt.Println(runCommand)

	const memoryLimit = 1024 * 1024 * 256 // 256 MB
	timeout := 5                          // 5 seconds

	// Create the container
	response, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image:           selectDockerImageForLanguage(language),
			Cmd:             runCommand,
			NetworkDisabled: true,
			WorkingDir:      "/usr/src/app/runs",
			Hostname:        runId,
		},
		&container.HostConfig{
			Binds: []string{
				filepath.Join(currentDir, "runs") + ":" + "/usr/src/app/runs",
			},
			RestartPolicy: container.RestartPolicy{
				Name: "no",
			},
			Resources: container.Resources{
				Memory: memoryLimit, // 256 MB
			},
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		panic(err)
	}

	// start the container, if it returns an error, print it
	if err := cli.ContainerStart(ctx, response.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// stop the container
	if err := cli.ContainerStop(ctx, response.ID, *&container.StopOptions{
		Signal:  "SIGTERM",
		Timeout: &timeout, // pass the address of the timeout variable
	}); err != nil {
		panic(err)
	}

	// get the logs from the container
	out, err := cli.ContainerLogs(ctx, response.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	}

	// clean up the container and the file
	defer CleanUpContainerAndFiles(cli, response, ctx, filePath)

	// ignore first 8 bits of nonsense
	ignore := make([]byte, 8)
	out.Read(ignore)

	// read the rest of the output
	output, err := io.ReadAll(out)
	if err != nil {
		panic(err)
	}

	return output
}

# Remote Code Execution (RCE) API

The RCE API is a HTTP-based service written in Go. It allows you to execute programs written in various languages remotely.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

-   Go (version 1.21.3 or later)
-   Docker

### Installation

1. Clone the repository to your local machine using `git clone`.
2. Navigate to the project directory.
3. Install the necessary Go packages using `go get`.
4. Build the project using `go build -tags netgo -ldflags '-s -w' -o rce`.

### Configuration

Create a `.env` file in the project directory with the following contents:

```bash
PORT=8000 # The port on which the server will run
```

### Pulling Docker Images

Before starting the server, you need to pull the Docker images for the different languages that the API supports. You can do this by running:

```bash
./rce pull-images
```

### Starting the Server

Start the server by running:

```bash
./rce serve
```

The server will start running at <http://localhost:$PORT>.

## Usage

To execute your code, make a POST request to <http://localhost:$PORT/run> with the following JSON body:

```json
{
    "code": "print(\"Hello World\")",
    "language": "python"
}
```

The API will execute the code and return a JSON object with the output:

```json
{
    "output": "Hello World"
}
```

## Building Docker Images

You can build the Docker images by running:

```bash
docker build -t huddl/rce .
```

You can run the Docker images by running:

```bash
docker run -it -p 8000:8000 -v "/var/run/docker.sock:/var/run/docker.sock" -v "/usr/src/app/runs:/usr/src/app/runs" huddl/rce
```

## Supported Languages

Currently, the API supports the following programming languages:

-   Python
-   C
-   C++
-   Java

We are actively working on adding support for more languages. Stay tuned!

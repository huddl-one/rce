#!/bin/bash

# Run the Python tests with resource limits and timeout
start_time=$(date +%s)  # Record start time

docker run --rm -v "$(pwd)/app.py:/app/app.py" -v "$(pwd)/test.py:/app/test.py" --name dynamic-python-executor --memory=2g --cpus=2 python:3.8-slim bash -c "\
python - <<END
import subprocess

try:
    subprocess.run(['timeout', '10s', 'python', '/app/test.py'], check=True)
except subprocess.TimeoutExpired:
    print('Execution timed out.')
END
"

end_time=$(date +%s)    # Record end time

duration=$((end_time - start_time))  # Calculate duration in seconds

echo "Total execution time: $duration seconds"

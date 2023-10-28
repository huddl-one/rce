FROM python:3.8-slim

# Set the working directory in the container
WORKDIR /app

# Copy the local code into the container at /app
COPY . /app

# Install the parameterized library
RUN pip3 install parameterized

# Define the default command to run when the container starts
CMD ["python", "app.py"]

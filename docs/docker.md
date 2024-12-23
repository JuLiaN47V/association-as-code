# Running the Docker Image

This guide explains how to run the Docker image for the project.

## Prerequisites

- Docker installed on your machine. You can download and install Docker from [here](https://www.docker.com/get-started).

## Building the Docker Image

To build the Docker image, navigate to the root directory of the project and run the following command:

```sh
docker build -t association-as-code .
```

This command will build the Docker image and tag it as `association-as-code`.

## Getting started

To run the Docker container, use the following command:

```sh
docker run -d -p 4000:80 --name association-as-code association-as-code
```

This command will run the Docker container in detached mode and map port `4000` of the host to port `80` of the container. The container will be named `association-as-code`.

### Accessing the Application

Once the container is running, you can access the application by navigating to `http://localhost:4000` in your web browser.

## Using a custom `config.yaml`

To use a custom `config.yaml` file with the Docker container, you can mount the file into the container. Use the following command to run the container with the `config.yaml` file mounted:

```sh
docker run -d -p 4000:80 --name association-as-code -v /path/to/your/config.yaml:/app/config.yaml association-as-code
```

Replace `/path/to/your/config.yaml` with the actual path to your `config.yaml` file. This command will mount the `config.yaml` file from your host machine into the container at `/app/config.yaml`.

## Adding static files with the `static` Directory

To add the `static` directory to the Docker container, you can mount the directory into the container. Use the following command to run the container with the `static` directory mounted:

```sh
docker run -d -p 4000:80 --name association-as-code -v /path/to/your/static:/app/static association-as-code
```

Replace `/path/to/your/static` with the actual path to your `static` directory. This command will mount the `static` directory from your host machine into the container at `/app/static`.  
The `static` directory has to include following subdirectories: `css`, `fonts` `img` and `js`.
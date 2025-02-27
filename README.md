# blyf

A basic fileserver for uploading and downloading small files.

## Attention!

This project is still work in progress!

## Testing

You can either install the go binary on your machine (https://go.dev/doc/install) or use the testing.sh script.

### Using the sript

Attention: The run_go.sh script needs podman to be installed on your machine and the go image needs to be downloaded.

#### Making the script executable

```bash
chmod +x run_go.sh
```

#### Running the script

The testing script allows you to directly use a docker image and run the application inside it.
This can be useful if you don't want to install the go binaries on your machine and you cannot build a docker image.
Depending weather you want to run directly the go file or you want to build you go project into a binary, run the following:

```bash
./run_go.sh go run ./cmd/server/main.go
```

```bash
./run_go.sh go build -o blyf ./cmd/server/
```

Or just use `./run_go.sh go` to run the go cli and see what commands are available.

### Using the docker image

You can run the application via docker.
First, you need to build the application:

```bash
# you can use either docker or podman
docker build -t blyf:dev . # or use whatever tag you prefer
```

Second, run the application:

```bash
docker run blyf:dev
```

Or you can run one of the dev shell scripts:

```bash
# depending on which container runtime you have, use either one
./dev_docker.sh
./dev_podman.sh
```

Double check if the files are executable. If not run:

```bash
chmod +x <script>
```

## Testing the file upload

When just using the terminal, you can use either curl or the e2e_testing.sh script.
When using the e2e_testing shell script, double check if it is executable, beforehand.
When running curl commands, you can do something like this:

```bash
curl -X POST -F "file=@ test.txt" http://localhost:8080/upload
```

Ultimately, the e2e_testing script does more or less the same.

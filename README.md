# blyf

## Testing

You can either install the go binary on your machine (https://go.dev/doc/install) or use the testing.sh script.

### Using the sript

Attention: The testing.sh script needs podman to be installed on your machine and the go image needs to be downloaded.

#### Making the script executable

```
chmod +x testing.sh
```

#### Running the script

Depending weather you want to run directly the go file or you want to build you go project into a binary, run the following:

```
./testing.sh go run main.go
```

```
./testing.sh go build -o bin/blyf
```

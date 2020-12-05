## commands
- `go run file_nmame` to run the code
- `go help` to get a list of other commands
- `go get` to download the packaged named by the import paths, w/ dependencies
    - installed in `$GOPATH/bin`
    - `/Users/ursname/go/bin`

## using external modules
- a `go.mod` file must exist
    - a way to keep track of `.go` files and its dependencies
- run: `go mod init hello`
    - hello = name of the module for the code
- the go tool with module support will then automagically download and install the dependencies
    - downloading the latest version by default

##
- `panic()` - for handling unexpected exceptions, prints goroutine traces, exit with non-zero status
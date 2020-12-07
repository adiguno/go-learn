## commands
- `go run file_nmame` to run the code
- `go help` to get a list of other commands
- `go get` to download the packaged named by the import paths, w/ dependencies
    - installed in `$GOPATH/bin`
    - `/Users/ursname/go/bin`

## file structure
 - A module = a collection of packages that are released, versioned, and distributed together. 
    - Modules may be downloaded directly from version control repositories or from module proxy servers.
    - identified by a module path, which is declared in a go.mod file, together with information about the module's dependencies. 
    - The module root directory = the directory that contains the go.mod file. 
    - The __main__ module = the module containing the directory where the go command is invoked.

 - Each package within a module is a collection of source files in the same directory that are compiled together.
    - A package path is the module path joined with the subdirectory containing the package (relative to the module root). 
        - e.g., the module "golang.org/x/net" contains a package in the directory "html". 
        - That package's path is "golang.org/x/net/html".



## using external modules
- a `go.mod` file must exist
    - a way to keep track of `.go` files and its dependencies
- run: `go mod init hello`
    - hello = name of the module for the code
- the go tool with module support will then automagically download and install the dependencies
    - downloading the latest version by default

##
- `panic()` - for handling unexpected exceptions, prints goroutine traces, exit with non-zero status
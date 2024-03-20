# file-server-go

## Setting up dir
- create dir and cd
- Module: Manage dependencies and vc
  - go mod init github.com/rsmadhavan/file-server-go.git
  - use the github url of repo without https://
- add external packages/dependencies using go get
- Build and run
  - go build followed by ./myproject &rarr; compiles to binary
  - go run main.go  &rarr; to directly run without compiling

## To check if a package is available in your workspace, you can use the go list command. Here's how:

```bash
go list -m -f {{.}} <package-import-path>
```
Explanation:

go list: This command provides information about Go packages.
-m: Specifies that the operation is performed on the current module and its dependencies.
-f: Specifies the format for printing the results. In this case, we use {{.}} to print just the package path.
Replace <package-import-path> with the import path of the package you want to check.

For example, if you want to check if the package github.com/gorilla/mux is available in your workspace, you would run:

bash
Copy code
go list -m -f {{.}} github.com/gorilla/mux
If the package is available, this command will print the import path of the package. If it's not available, there will be no output.

Make sure you are running this command in the root directory of your Go module (if you're using modules), or inside your $GOPATH/src directory (if you're not using modules).

## tests
Run tests using go test ./...





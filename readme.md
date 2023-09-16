### Basic Go Rest APIs app

1. Install Go
2. To create a Go application - run `go mod init {app-name}`
3. Create a file called `main.go`
4. Create the data types using `struct`
5. Create handler functions for each type of request - fetch all, fetch by key, create new
6. Write `main` function with a router(in this case, we use `gin`)
7. Create router endpoints based on HTTP methods and link them with the appropriate request handler functions
8. Run `go mod tidy` to install all dependencies/packages
9. Run `go run .` to start the app server
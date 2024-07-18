# Testing Go as server and api

To run this:

```bash
go mod tidy
go run .
```

## Modules

When you create a new module:

- run `go mod init <package-name>`: this creates a "go.mod" file,
- create a <something.go> file, and declare `package <package-name>`,
- any function with CapitalLetter is public

You can use it in `main` with `import`.
In "go.mod" file, the "import" should have the line:

```go
replace server => ./server
```

and run `go mod tidy`.

## Examples of Server and API call

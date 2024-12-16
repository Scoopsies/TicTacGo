# TicTacGo


### Testing

To run all tests once:
```
ginkgo ./...
```

To continually run tests while making changes:
```
ginkgo watch -r
```

To show test coverage:
```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out   
```

A browser will open with a drop-down list off all packages and coverage.
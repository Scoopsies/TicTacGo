# TicTacGo

Tic Tac Go is a Tic Tac Toe game created in GO over the course of 5 days. It its played in the terminal.

It can be played:
- Human vs Human
- Human vs Computer
- Computer vs Human
- Computer vs Human

It has three difficulty levels for the computer:
- Easy
- Medium
- Hard

---

### How to install and play:

Start by verifying that you have go installed:
```
go version
```

If not it can be installed [here](https://go.dev/doc/install).

Clone this repo:
```
git clone git@github.com:Scoopsies/TicTacGo.git
```

Run the program:
```
cd TicTacGo
go run .
```
---
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
test:
	mkdir -p .testoutput
	go test -v -coverprofile=.testoutput/coverage.out ./...
	go tool cover -func=.testoutput/coverage.out

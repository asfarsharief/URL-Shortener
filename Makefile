#Makefile to run the application, run test cases, and build docker image 
#=======================================================================

run:
	@go run server/main.go

test:
	@go test -coverprofile=.code_coverage.out ./...

show-code-coverage: test
	@mkdir -p .code_coverage
	@go tool cover -html=.code_coverage.out -o .code_coverage/index.html
	@cd .code_coverage && python -m SimpleHTTPServer 3001
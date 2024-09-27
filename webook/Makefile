# command: make clean docker
.PHONY: clean docker
docker:
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -o webook .
	@docker build -t flycash/webook:v0.0.1 .

 .PHONY: clean
 clean:
	@rm webook || true
	@docker rmi -f flycash/webook:v0.0.1

build:
	@rm -f function.zip
	@GOOS=linux GOARCH=amd64 go build -o bootstrap
	@powershell -Command "Compress-Archive bootstrap function.zip"
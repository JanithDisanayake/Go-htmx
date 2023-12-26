run:
	@bash init.sh
	@echo " server is up and running 🚀"
	@go mod tidy
	@go run main.go
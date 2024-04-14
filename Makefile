run: build
	@./bin/redish
	
build:
	@go build -o bin/redish .

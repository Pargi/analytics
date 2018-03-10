build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/health health/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/events events/post.go

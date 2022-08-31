run-dev:
	APP_ENV=DEVELOPMENT go run ./cmd/playlist-http/

build:
	go build -v -o bin/playlist-http cmd/playlist-http/*.go
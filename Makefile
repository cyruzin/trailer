build:
	go build -ldflags "-s -w" -ldflags "-X cmd.version=1.0.1" -o "dist/trailer"
build:
	go build -ldflags "-s -w" -ldflags "-X cmd.version=1.1.3" -o "dist/trailer"
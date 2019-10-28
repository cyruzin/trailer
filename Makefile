build:
	go build -ldflags "-s -w" -ldflags "-X cmd.version=1.1.2" -o "dist/trailer"
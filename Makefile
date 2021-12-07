version=1.2.0

build:
	go build -ldflags "-X 'github.com/cyruzin/trailer/cmd/trailer.TrailerVersion=${version}'" -o "dist/trailer"

build-win:
	go build -ldflags "-X 'github.com/cyruzin/trailer/cmd/trailer.TrailerVersion=${version}'" -o "dist/trailer.exe"
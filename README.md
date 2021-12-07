# Trailer

[![build](https://github.com/cyruzin/trailer/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/cyruzin/trailer/actions/workflows/build.yml)

Trailer is a cli tool that will quickly bring the trailers of any movie or tv show with a few commands.

## Build

Download needed packages.

```sh
go mod download
```

Type the command below in the terminal to create the binary.

```sh
make build
```

### Linux

Once the build is finished, move the binary to **/usr/bin** folder.

```sh
sudo mv ./trailer /usr/bin
```

## Usage

For Movies:

```sh
trailer movie john wick parabellum
```

For TV Shows:

```sh
trailer tv game of thrones
```

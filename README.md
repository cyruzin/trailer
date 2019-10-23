# Trailer

[![Build Status](https://travis-ci.org/cyruzin/trailer.svg?branch=master)](https://travis-ci.org/cyruzin/trailer)

Trailer is a cli tool that will quickly bring the trailers of any movie or tv show with a few commands.

## Build

Download needed packages.

```sh
go mod download
```

Rename the environment file example and add your own TMDb API Key.

```sh
cp .env.example .env
```

Or set a environment variable:

```sh
export TMDB_KEY=YOUR_TMDB_KEY
```

Type the command below in the terminal to create the binary.

```sh
make build
```

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

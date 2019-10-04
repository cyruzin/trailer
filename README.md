# Trailer

[![Build Status](https://travis-ci.org/cyruzin/trailer.svg?branch=master)](https://travis-ci.org/cyruzin/trailer)

Trailer is a cli tool that will quickly bring the trailers of any movie or tv show with a few commands.

## Build

Download needed packages
```sh
go get
```
Copy environment file example to your current, add your own TMDB Api Key

```sh
cp .env.example .env
```
Type the command below in the terminal to create the binary.

```sh
make build
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

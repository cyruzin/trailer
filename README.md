# Trailer

[![build](https://github.com/cyruzin/trailer/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/cyruzin/trailer/actions/workflows/build.yml) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

Trailer is a cli tool that will quickly bring the trailers of any movie or tv show with a few commands.

## Download

Head to releases page and download the binary that suits you best.

## Usage

For Movies:

```sh
trailer movie john wick parabellum
```

For TV Shows:

```sh
trailer tv game of thrones
```

## Usage With Language Flag

For Movies:

```sh
trailer movie john wick parabellum --lang=pt-BR
```

For TV Shows:

```sh
trailer tv game of thrones --lang=pt-BR
```

## Build

Download needed packages. Type:

```sh
go mod download
```

Type the command below in the terminal to create the binary:

For Linux / Mac

```sh
make build
```

For Windows

```sh
make build-win
```

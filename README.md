[![GoReportCard](http://goreportcard.com/badge/github.com/bubbajoe/avgo)](http://goreportcard.com/report/github.com/bubbajoe/avgo)
[![Go Reference](https://pkg.go.dev/badge/github.com/bubbajoe/avgo.svg)](https://pkg.go.dev/github.com/bubbajoe/avgo)
[![Coveralls](https://coveralls.io/repos/github/bubbajoe/avgo/badge.svg?branch=master)](https://coveralls.io/github/bubbajoe/avgo)

`avgo` is a Golang library providing C bindings for [ffmpeg](https://github.com/FFmpeg/FFmpeg)

forked from [bubbajoe/avgo](https://github.com/bubbajoe/go-avgo)

It's only compatible with `ffmpeg` `n4.4`.

Its main goals are to:
- [x] provide a better GO idiomatic API
    - typed constants and flags
    - standard error pattern
    - struct-based functions
    - use of common go interfaces
- [x] provide the GO version of [ffmpeg examples](https://github.com/FFmpeg/FFmpeg/tree/n4.4/doc/examples)
- [x] be fully tested
- [x] be fully documented
- [x] be fully compatible with all `ffmpeg` libs

# Examples

Examples are located in the [examples](examples) directory and mirror as much as possible the [ffmpeg examples](https://github.com/FFmpeg/FFmpeg/tree/n4.4/doc/examples).

|name|avgo|ffmpeg|
|---|---|---|
|Demuxing/Decoding|[see](examples/demuxing_decoding/main.go)|[see](https://github.com/FFmpeg/FFmpeg/blob/n4.4/doc/examples/demuxing_decoding.c)
|Filtering|[see](examples/filtering/main.go)|[see](https://github.com/FFmpeg/FFmpeg/blob/n4.4/doc/examples/filtering_video.c)
|Remuxing|[see](examples/remuxing/main.go)|[see](https://github.com/FFmpeg/FFmpeg/blob/n4.4/doc/examples/remuxing.c)
|Transcoding|[see](examples/transcoding/main.go)|[see](https://github.com/FFmpeg/FFmpeg/blob/n4.4/doc/examples/transcoding.c)
|Resampling|TODO|TODO
|AVIO|TODO|TODO
|GOCV|TODO|TODO

*Tip: you can use the video sample located in the `testdata` directory for your tests*

# Install ffmpeg from source

If you don't know how to install `ffmpeg`, you can use the following to install it from source:

```sh
$ make install-ffmpeg
```

`ffmpeg` will be built from source in a directory named `tmp` and located in you working directory

For your GO code to pick up `ffmpeg` dependency automatically, you'll need to add the following environment variables:

(don't forget to replace `{{ path to your working directory }}` with the absolute path to your working directory)

```sh
export CGO_LDFLAGS="-L{{ path to your working directory }}/tmp/n4.4/lib/",
export CGO_CXXFLAGS="-I{{ path to your working directory }}/tmp/n4.4/include/",
export CGO_CFLAGS="-I{{ path to your working directory }}/tmp/n4.4/include/",
export PKG_CONFIG_PATH="{{ path to your working directory }}/tmp/n4.4/lib/pkgconfig",
```
# go-imgdye

Package **imgdye** provides tools dying an image, for the Go programming language.

One use case for this is giving a (new) image a background-color.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-imgdye

[![GoDoc](https://godoc.org/github.com/reiver/go-imgdye?status.svg)](https://godoc.org/github.com/reiver/go-imgdye)

## Examples

Here is an example of making a whole image one single single.

```golang
import "github.com/reiver/go-imgdye"

// ...

var img draw.Image // = ...

// ...

var c color.Color = color.NRGBA{0xFF,0x1D,0xCE, 0xFF} // #FF1DCE i.e., hot magenta

imgdye.Dye(img, c)
```

## Import

To import package **imgdye** use `import` code like the follownig:
```
import "github.com/reiver/go-imgdye"
```

## Installation

To install package **imgdye** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-imgdye
```

## Author

Package **imgdye** was written by [Charles Iliya Krempeaux](http://changelog.ca)

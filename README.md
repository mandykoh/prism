# prism

[![GoDoc](https://godoc.org/github.com/mandykoh/prism?status.svg)](https://godoc.org/github.com/mandykoh/prism)
[![Go Report Card](https://goreportcard.com/badge/github.com/mandykoh/prism)](https://goreportcard.com/report/github.com/mandykoh/prism)
[![Build Status](https://travis-ci.org/mandykoh/prism.svg?branch=main)](https://travis-ci.org/mandykoh/prism)

`prism` aims to become a set of utilities for practical colour management and conversion.

`prism` currently implements encoding/decoding linear colour from sRGB and Adobe RGB with optional fast LUT-based conversion, and conversion to and from CIE XYZ and CIE Lab.

See the [API documentation](https://godoc.org/github.com/mandykoh/prism) for more details.

This software is made available under an [MIT license](LICENSE).


## Example usage

Image data provided by the standard [`image`](https://golang.org/pkg/image/) package doesn’t come with colour profile information. However, interpreting the image data directly as raw, linear RGB values for image processing purposes is unlikely to produce good or correct results as nearly all images are encoded with non-linear values referencing specific colour spaces.

`prism` can be used to convert between encoded colour values and a normalised, linear representation more suitable for image processing, and subsequently converting back to encoded colour values in (potentially) other colour spaces.

The following example converts Adobe RGB (1998) pixel data to sRGB. It retrieves a pixel from an [NRGBA image](https://golang.org/pkg/image/#NRGBA), decodes it as an Adobe RGB (1998) linearised colour value, then converts that to an sRGB colour value via the CIE XYZ intermediate colour space, before finally encoding the result as an 8-bit sRGB value suitable for writing back to an `image.NRGBA`:

```go
c := inputImg.NRGBAAt(x, y)             // Take input colour value
ac := adobergb.ColorFromNRGBA(c)        // Interpret image pixel as Adobe RGB and convert to linear representation
sc := srgb.ColorFromXYZ(ac.ToXYZ())     // Convert to XYZ, then from XYZ to sRGB linear representation
outputImg.SetNRGBA(x, y, sc.ToNRGBA())  // Write sRGB-encoded value to output image
``` 

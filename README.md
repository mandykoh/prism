# prism

[![GoDoc](https://godoc.org/github.com/mandykoh/prism?status.svg)](https://godoc.org/github.com/mandykoh/prism)
[![Go Report Card](https://goreportcard.com/badge/github.com/mandykoh/prism)](https://goreportcard.com/report/github.com/mandykoh/prism)
[![Build Status](https://travis-ci.org/mandykoh/prism.svg?branch=main)](https://travis-ci.org/mandykoh/prism)

`prism` aims to become a set of utilities for practical colour management and conversion in pure Go.

`prism` currently implements encoding/decoding linear colour from sRGB, Adobe RGB, Pro Photo RGB, and Display P3 encodings with optional fast LUT-based conversion, and conversion to and from CIE xyY, CIE XYZ, and CIE Lab. Chromatic adaptation in XYZ space between different white points is also supported.

`prism` does not yet support detection/extraction/embedding of tagged colour profiles in images, conversions between arbitrary ICC profiles, nor CMYK.

See the [API documentation](https://godoc.org/github.com/mandykoh/prism) for more details.

This software is made available under an [MIT license](LICENSE).

Much of this implementation is based on information provided by [Bruce Lindbloom](http://www.brucelindbloom.com) and [Charles Poynton](http://poynton.ca), among many others who generously contribute to public edification on the esoteric science of colour.


## Example usage

Image data provided by the standard [`image`](https://golang.org/pkg/image/) package doesn’t come with colour profile information. However, interpreting the image data directly as raw, linear RGB values for image processing purposes is unlikely to produce good or correct results as nearly all images are encoded with non-linear values referencing specific colour spaces. Poorly converted images can themselves result in loss or exaggeration of saturation and contrast, colour shifts, and loss of fidelity, while using such images for image processing will result in incorrect blending, sharpening, blurring, etc.

`prism` can be used to convert between encoded colour values and a normalised, linear representation more suitable for image processing, and subsequently converting back to encoded colour values in (potentially) other colour spaces.


### Colour conversion

Conversions between RGB colour spaces are performed via the CIE XYZ intermediate colour space (using the `ToXYZ` and `ColorFromXYZ` functions).

The following example converts Adobe RGB (1998) pixel data to sRGB. It retrieves a pixel from an [NRGBA image](https://golang.org/pkg/image/#NRGBA), decodes it as an Adobe RGB (1998) linearised colour value, then converts that to an sRGB colour value via CIE XYZ, before finally encoding the result as an 8-bit sRGB value suitable for writing back to an `image.NRGBA`:

```go
c := inputImg.NRGBAAt(x, y)                 // Take input colour value
ac, alpha := adobergb.ColorFromNRGBA(c)     // Interpret image pixel as Adobe RGB and convert to linear representation
sc := srgb.ColorFromXYZ(ac.ToXYZ())         // Convert to XYZ, then from XYZ to sRGB linear representation
outputImg.SetNRGBA(x, y, sc.ToNRGBA(alpha)) // Write sRGB-encoded value to output image
``` 


### Chromatic adaptation

Adobe RGB (1998) and sRGB are both specified referring to a standard D65 white point. However, Pro Photo RGB references a D50 white point. When converting between white points, a chromatic adaptation is required to compensate for a shift in warmness/coolness that would otherwise occur.

The following example prepares such a chromatic adaptation (using the `AdaptBetweenWhitePoints` function), then uses it in converting from Pro Photo RGB to sRGB:

```go
adaptation := ciexyz.AdaptBetweenWhitePoints(
    prophotorgb.StandardWhitePoint,         // From D50
    srgb.StandardWhitePoint,                // To D65
)

c := inputImg.NRGBAAt(x, y)                 // Take input colour value
pc, alpha := prophotorgb.ColorFromNRGBA(c)  // Interpret image pixel as Pro Photo RGB and convert to linear representation

xyz := pc.ToXYZ()                           // Convert from Pro Photo RGB to CIE XYZ
xyz = adaptation.Apply(xyz)                 // Apply chromatic adaptation from D50 to D65

sc := srgb.ColorFromXYZ(xyz)                // Convert from CIE XYZ to sRGB linear representation
outputImg.SetNRGBA(x, y, sc.ToNRGBA(alpha)) // Write sRGB-encoded value to output image
```

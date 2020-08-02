package jpegmeta

import "fmt"

type markerType int

const (
	MarkerTypeInvalid                 markerType = 0x00
	MarkerTypeStartOfFrameBaseline    markerType = 0xc0
	MarkerTypeStartOfFrameProgressive markerType = 0xc2
	MarkerTypeDefineHuffmanTable      markerType = 0xc4
	MarkerTypeRestart0                markerType = 0xd0
	MarkerTypeRestart1                markerType = 0xd1
	MarkerTypeRestart2                markerType = 0xd2
	MarkerTypeRestart3                markerType = 0xd3
	MarkerTypeRestart4                markerType = 0xd4
	MarkerTypeRestart5                markerType = 0xd5
	MarkerTypeRestart6                markerType = 0xd6
	MarkerTypeRestart7                markerType = 0xd7
	MarkerTypeStartOfImage            markerType = 0xd8
	MarkerTypeEndOfImage              markerType = 0xd9
	MarkerTypeStartOfScan             markerType = 0xda
	MarkerTypeDefineQuantisationTable markerType = 0xdb
	MarkerTypeDefineRestartInterval   markerType = 0xdd
	MarkerTypeApp0                    markerType = 0xe0
	MarkerTypeApp1                    markerType = 0xe1
	MarkerTypeApp2                    markerType = 0xe2
	MarkerTypeApp3                    markerType = 0xe3
	MarkerTypeApp4                    markerType = 0xe4
	MarkerTypeApp5                    markerType = 0xe5
	MarkerTypeApp6                    markerType = 0xe6
	MarkerTypeApp7                    markerType = 0xe7
	MarkerTypeApp8                    markerType = 0xe8
	MarkerTypeApp9                    markerType = 0xe9
	MarkerTypeApp10                   markerType = 0xea
	MarkerTypeApp11                   markerType = 0xeb
	MarkerTypeApp12                   markerType = 0xec
	MarkerTypeApp13                   markerType = 0xed
	MarkerTypeApp14                   markerType = 0xee
	MarkerTypeApp15                   markerType = 0xef
	MarkerTypeComment                 markerType = 0xfe
)

func (mt markerType) String() string {
	switch mt {
	case MarkerTypeStartOfFrameBaseline:
		return "SOF0"
	case MarkerTypeStartOfFrameProgressive:
		return "SOF2"
	case MarkerTypeDefineHuffmanTable:
		return "DHT"
	case MarkerTypeRestart0:
		return "RST0"
	case MarkerTypeRestart1:
		return "RST1"
	case MarkerTypeRestart2:
		return "RST2"
	case MarkerTypeRestart3:
		return "RST3"
	case MarkerTypeRestart4:
		return "RST4"
	case MarkerTypeRestart5:
		return "RST5"
	case MarkerTypeRestart6:
		return "RST6"
	case MarkerTypeRestart7:
		return "RST7"
	case MarkerTypeStartOfImage:
		return "SOI"
	case MarkerTypeEndOfImage:
		return "EOI"
	case MarkerTypeStartOfScan:
		return "SOS"
	case MarkerTypeDefineQuantisationTable:
		return "DQT"
	case MarkerTypeDefineRestartInterval:
		return "DRI"
	case MarkerTypeApp0:
		return "APP0"
	case MarkerTypeApp1:
		return "APP1"
	case MarkerTypeApp2:
		return "APP2"
	case MarkerTypeApp3:
		return "APP3"
	case MarkerTypeApp4:
		return "APP4"
	case MarkerTypeApp5:
		return "APP5"
	case MarkerTypeApp6:
		return "APP6"
	case MarkerTypeApp7:
		return "APP7"
	case MarkerTypeApp8:
		return "APP8"
	case MarkerTypeApp9:
		return "APP9"
	case MarkerTypeApp10:
		return "APP10"
	case MarkerTypeApp11:
		return "APP11"
	case MarkerTypeApp12:
		return "APP12"
	case MarkerTypeApp13:
		return "APP13"
	case MarkerTypeApp14:
		return "APP14"
	case MarkerTypeApp15:
		return "APP15"
	case MarkerTypeComment:
		return "COM"
	default:
		return fmt.Sprintf("Unknown (%0x)", byte(mt))
	}
}

/*
Package tutorial is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /home/raphael/go/src/github.com/openconfig/ygot/genutil/names.go
using the following YANG input files:
	- demo-port.yang
Imported modules were sourced from:
*/
package tutorial

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
)


var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x59, 0x5f, 0x6f, 0xda, 0x30,
		0x10, 0x7f, 0xe7, 0x53, 0x44, 0x7e, 0xa6, 0x02, 0x3a, 0x6d, 0x0f, 0xbc, 0xc1, 0xa0, 0xdb, 0xd4,
		0xb1, 0xa1, 0x32, 0xed, 0x65, 0xaa, 0x90, 0x4b, 0x0e, 0x6a, 0x8d, 0xd8, 0xc8, 0x71, 0x34, 0x50,
		0xc5, 0x77, 0x9f, 0x12, 0x27, 0x2c, 0x09, 0x49, 0x7c, 0x36, 0xb4, 0x53, 0x2b, 0xfa, 0x10, 0x35,
		0xce, 0x9d, 0x7d, 0x77, 0xbf, 0x1f, 0xf7, 0x27, 0x79, 0x6a, 0x79, 0x9e, 0xe7, 0x91, 0x6f, 0x34,
		0x00, 0xd2, 0xf7, 0x08, 0x69, 0xeb, 0xfb, 0x5b, 0xc6, 0x7d, 0xd2, 0xf7, 0xba, 0xe9, 0xed, 0x47,
		0xc1, 0x97, 0x6c, 0x95, 0x5b, 0x18, 0x31, 0x49, 0xfa, 0x9e, 0x56, 0x4e, 0x16, 0x36, 0x42, 0xaa,
		0xb0, 0xb0, 0x54, 0xd8, 0x57, 0x3f, 0x6e, 0x17, 0x1f, 0xa6, 0x87, 0xf4, 0x4a, 0xcb, 0xe5, 0xc3,
		0x0e, 0x0f, 0xa6, 0x12, 0x96, 0x6c, 0x7b, 0x74, 0x48, 0xe1, 0x20, 0x1f, 0x02, 0x71, 0x15, 0x9f,
		0x56, 0x3a, 0x2c, 0x11, 0x9a, 0x89, 0x48, 0x2e, 0xa0, 0x72, 0x03, 0x6d, 0x10, 0xec, 0xfe, 0x08,
		0xe9, 0x27, 0x06, 0xeb, 0xb3, 0xda, 0xd5, 0x82, 0x9f, 0x69, 0x38, 0x90, 0xab, 0x28, 0x00, 0xae,
		0x48, 0xdf, 0x53, 0x32, 0x82, 0x1a, 0xc1, 0x9c, 0x54, 0xce, 0xb4, 0x23, 0xd9, 0x7d, 0x61, 0x65,
		0x5f, 0xf2, 0xbb, 0x1c, 0xec, 0x42, 0xd0, 0xeb, 0xbd, 0xc9, 0xc7, 0xbe, 0xce, 0x91, 0x6a, 0x08,
		0x8c, 0x50, 0x60, 0x20, 0xb1, 0x82, 0x06, 0x0b, 0x91, 0x35, 0x54, 0xd6, 0x90, 0xd9, 0x42, 0x57,
		0x0d, 0x61, 0x0d, 0x94, 0x46, 0x48, 0x0f, 0x02, 0x8b, 0x2c, 0xf2, 0x86, 0x38, 0x64, 0xc1, 0x4d,
		0xe5, 0x0d, 0x3e, 0x35, 0xc3, 0x8d, 0x86, 0xdd, 0x06, 0x7e, 0x27, 0x1a, 0xd8, 0xd2, 0xc1, 0x99,
		0x16, 0xce, 0xf4, 0x70, 0xa5, 0x49, 0x33, 0x5d, 0x0c, 0xb4, 0x41, 0xd3, 0xe7, 0x20, 0x18, 0x6e,
		0x00, 0x7c, 0x7c, 0xf8, 0x32, 0x8c, 0xb4, 0x1a, 0x32, 0x02, 0xc5, 0x5a, 0x61, 0x14, 0xc7, 0x92,
		0xcb, 0x85, 0x64, 0x27, 0x91, 0xcd, 0x95, 0x74, 0x27, 0x93, 0xef, 0x64, 0x12, 0x9e, 0x4a, 0x46,
		0x1c, 0x29, 0x91, 0xe4, 0x3c, 0x18, 0xf3, 0x63, 0xb7, 0x01, 0x37, 0xd4, 0x98, 0x0f, 0x5c, 0x31,
		0xb5, 0x93, 0xb0, 0xb4, 0xc1, 0x2d, 0xcb, 0x6e, 0xef, 0x2d, 0x74, 0xbe, 0xa4, 0x47, 0x0d, 0x69,
		0xe8, 0x80, 0x78, 0x66, 0xf0, 0x6c, 0x3a, 0x1e, 0x8f, 0x6c, 0xd1, 0xfe, 0x49, 0xd7, 0x11, 0xc4,
		0x3d, 0xd3, 0x2f, 0x2b, 0xbd, 0xf8, 0xef, 0xc9, 0x5a, 0xe3, 0xd8, 0xdc, 0x79, 0xaf, 0xfb, 0x69,
		0x48, 0xac, 0x37, 0xda, 0x5b, 0x69, 0xdc, 0x9f, 0x9b, 0x7b, 0xcf, 0x9c, 0x58, 0x07, 0x9c, 0x0b,
		0x45, 0x15, 0x13, 0x1c, 0x99, 0x5f, 0x17, 0x8f, 0x10, 0xd0, 0x0d, 0x55, 0x8f, 0x71, 0x5c, 0x3b,
		0x87, 0xdf, 0x5c, 0x27, 0x69, 0x78, 0x93, 0x6b, 0x07, 0x55, 0x9a, 0xf5, 0x6e, 0x4a, 0x46, 0x0b,
		0xc5, 0x53, 0x94, 0x46, 0x10, 0x88, 0xa9, 0x90, 0x6a, 0x1e, 0x5f, 0xc2, 0xe4, 0x3a, 0x4f, 0x93,
		0x67, 0xcb, 0x2d, 0x02, 0x0d, 0xde, 0x27, 0x5d, 0xe2, 0x15, 0x8f, 0x82, 0x07, 0x90, 0xf8, 0xae,
		0x23, 0xaf, 0x84, 0x6b, 0x3d, 0xba, 0x97, 0xd6, 0xe3, 0x75, 0xb7, 0x1e, 0xe8, 0xac, 0xee, 0x40,
		0x92, 0x3c, 0x51, 0x3e, 0x20, 0x44, 0xef, 0x28, 0x5f, 0x01, 0x3a, 0x81, 0x5a, 0x14, 0xa2, 0x09,
		0xe3, 0x0e, 0xe5, 0xdf, 0xaa, 0x0f, 0x2a, 0x96, 0x01, 0x73, 0x53, 0x7e, 0xa4, 0x77, 0x23, 0xe9,
		0x22, 0x4e, 0x54, 0x23, 0xb6, 0x62, 0xc9, 0xe8, 0xdd, 0xc5, 0xe7, 0x5a, 0x8b, 0x02, 0x39, 0xa1,
		0xdb, 0x17, 0x0f, 0xc5, 0xbb, 0xeb, 0x17, 0x8c, 0xc5, 0x99, 0xea, 0xce, 0xfd, 0x33, 0xe4, 0xe4,
		0x50, 0x51, 0x05, 0xf8, 0x6c, 0xac, 0xc5, 0xcf, 0x3c, 0x02, 0x5e, 0x5f, 0xf2, 0xf0, 0x2b, 0x1f,
		0x01, 0x15, 0x55, 0x51, 0xe8, 0x30, 0x03, 0x6a, 0xbd, 0xcb, 0x10, 0x78, 0x19, 0x02, 0xff, 0xcb,
		0x10, 0xf8, 0x20, 0xc4, 0x1a, 0x28, 0x77, 0x19, 0x00, 0x7b, 0x6f, 0x76, 0x96, 0xc0, 0xa4, 0x78,
		0xe4, 0x28, 0x31, 0x4b, 0xb6, 0x72, 0xad, 0x5a, 0x56, 0x6f, 0x3b, 0x6f, 0x61, 0x87, 0x6b, 0x04,
		0xc9, 0x57, 0x16, 0xaa, 0x81, 0x52, 0x86, 0x97, 0xa3, 0x13, 0xc6, 0xc7, 0x6b, 0x88, 0x39, 0x1e,
		0xe7, 0x35, 0x1e, 0xad, 0xd7, 0x0d, 0x45, 0x74, 0x42, 0xb7, 0x78, 0xe1, 0xef, 0xd2, 0x07, 0x09,
		0xfe, 0x70, 0x97, 0x8a, 0x5a, 0xb9, 0x89, 0x04, 0x1c, 0x03, 0x34, 0x69, 0xec, 0x0a, 0x0c, 0xd0,
		0x56, 0x83, 0xba, 0xb7, 0xfc, 0xea, 0x60, 0xf0, 0xc6, 0xe0, 0x45, 0xd5, 0xe7, 0x97, 0x26, 0xc3,
		0x8b, 0x36, 0xff, 0xb3, 0x4c, 0xff, 0x97, 0xda, 0x56, 0x67, 0x13, 0x61, 0xe1, 0x0d, 0xfd, 0x0d,
		0x77, 0x42, 0x64, 0x29, 0x53, 0x6b, 0xb5, 0xf6, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01, 0x00,
		0x00, 0xff, 0xff, 0x7d, 0x35, 0xd7, 0xea, 0xdf, 0x1a, 0x00, 0x00,
	}
)


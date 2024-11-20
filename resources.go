package main

import (
	"fmt"
)

var resData = map[string][]byte{
    
	"pkg.go.tpl": []byte{ 
		 112,  97,  99, 107,  97, 103, 101,  32, 123, 123,  32,  46,
		  67, 111, 110, 102, 105, 103,  46,  80,  97,  99, 107,  97,
		 103, 101,  32, 125, 125,  10,  10, 105, 109, 112, 111, 114,
		 116,  32,  40,  10,   9,  34, 102, 109, 116,  34,  10,  41,
		  10,  10, 118,  97, 114,  32, 114, 101, 115,  68,  97, 116,
		  97,  32,  61,  32, 109,  97, 112,  91, 115, 116, 114, 105,
		 110, 103,  93,  91,  93,  98, 121, 116, 101, 123,  10,  32,
		  32,  32,  32, 123, 123,  32, 114,  97, 110, 103, 101,  32,
		  46,  70, 105, 101, 108, 100, 115,  32, 125, 125,  10,   9,
		  34, 123, 123,  32,  46,  78,  97, 109, 101,  32, 125, 125,
		  34,  58,  32,  91,  93,  98, 121, 116, 101, 123,  32, 123,
		 123,  34,  92, 110,  92, 116,  92, 116,  34, 125, 125,  10,
		   9,  32,  32,  32,  32, 123, 123,  45,  32, 114,  97, 110,
		 103, 101,  32,  36, 105,  44,  32,  36, 118,  32,  58,  61,
		  32,  46,  66,  86,  97, 108, 117, 101,  32, 125, 125,  10,
		   9,   9, 123, 123,  45,  32,  36, 118,  32, 124,  32, 112,
		 114, 105, 110, 116, 102,  32,  34,  37,  52, 100,  34,  32,
		 125, 125,  44, 123, 123,  32, 105, 102,  32,  40, 105, 115,
		  87, 114,  97, 112,  32,  36, 105,  41,  32, 125, 125, 123,
		 123,  34,  92, 110,  92, 116,  92, 116,  34, 125, 125, 123,
		 123, 101, 110, 100,  32,  45, 125, 125,  10,   9,  32,  32,
		  32,  32, 123, 123,  32, 101, 110, 100,  32, 125, 125,  10,
		   9, 125,  44,  10,  32,  32,  32,  32, 123, 123,  32, 101,
		 110, 100,  32, 125, 125,  10, 125,  10,  10,  47,  47,  32,
		  82,  32,  45,  32, 114, 101, 116, 117, 114, 110,  32, 114,
		 101, 115, 111, 117, 114,  99, 101,  32,  40,  97, 115,  32,
		  98, 108, 111,  98,  41,  32,  98, 121,  32, 110,  97, 109,
		 101,  10, 102, 117, 110,  99,  32,  82,  40, 110,  97, 109,
		 101,  32, 115, 116, 114, 105, 110, 103,  41,  32,  40,  91,
		  93,  98, 121, 116, 101, 123, 123,  32, 105, 102,  32,  46,
		  67, 111, 110, 102, 105, 103,  46,  69, 114, 114, 111, 114,
		  32, 125, 125,  44,  32, 101, 114, 114, 111, 114, 123, 123,
		 101, 110, 100, 125, 125,  41,  32, 123,  10,  10,   9, 100,
		  97, 116,  97,  44,  32, 111, 107,  32,  58,  61,  32, 114,
		 101, 115,  68,  97, 116,  97,  91, 110,  97, 109, 101,  93,
		  10,  10,   9, 105, 102,  32,  33, 111, 107,  32, 123,  10,
		   9,  32,  32,  32,  32, 123, 123,  32, 105, 102,  32,  46,
		  67, 111, 110, 102, 105, 103,  46,  69, 114, 114, 111, 114,
		  32, 125, 125,  10,   9,   9, 114, 101, 116, 117, 114, 110,
		  32, 110, 105, 108,  44,  32, 102, 109, 116,  46,  69, 114,
		 114, 111, 114, 102,  40,  34,  85, 110, 107, 110, 111, 119,
		 110,  32, 114, 101, 115, 111, 117, 114,  99, 101,  58,  32,
		  37, 115,  34,  44,  32, 110,  97, 109, 101,  41,  10,   9,
		  32,  32,  32,  32, 123, 123,  32, 101, 108, 115, 101,  32,
		 125, 125,  10,   9,   9, 112,  97, 110, 105,  99,  40, 102,
		 109, 116,  46,  83, 112, 114, 105, 110, 116, 102,  40,  34,
		  85, 110, 107, 110, 111, 119, 110,  32, 114, 101, 115, 111,
		 117, 114,  99, 101,  58,  32,  37, 115,  34,  44,  32, 110,
		  97, 109, 101,  41,  41,  10,   9,  32,  32,  32,  32, 123,
		 123,  32, 101, 110, 100,  32, 125, 125,  10,   9, 125,  10,
		   9,  10,   9, 114, 101, 116, 117, 114, 110,  32, 100,  97,
		 116,  97, 123, 123,  32, 105, 102,  32,  46,  67, 111, 110,
		 102, 105, 103,  46,  69, 114, 114, 111, 114,  32, 125, 125,
		  44,  32, 110, 105, 108, 123, 123, 101, 110, 100, 125, 125,
		  10, 125,  10,  10,  47,  47,  32,  82, 115,  32,  45,  32,
		 114, 101, 116, 117, 114, 110,  32, 114, 101, 115, 111, 117,
		 114,  99, 101,  32,  40,  97, 115,  32, 115, 116, 114, 105,
		 110, 103,  41,  32,  98, 121,  32, 110,  97, 109, 101,  10,
		 102, 117, 110,  99,  32,  82, 115,  40, 110,  97, 109, 101,
		  32, 115, 116, 114, 105, 110, 103,  41,  32,  40, 115, 116,
		 114, 105, 110, 103, 123, 123,  32, 105, 102,  32,  46,  67,
		 111, 110, 102, 105, 103,  46,  69, 114, 114, 111, 114,  32,
		 125, 125,  44,  32, 101, 114, 114, 111, 114, 123, 123, 101,
		 110, 100, 125, 125,  41,  32, 123,  10,  32,  32,  32,  32,
		 123, 123,  32, 105, 102,  32,  46,  67, 111, 110, 102, 105,
		 103,  46,  69, 114, 114, 111, 114,  32, 125, 125,  10,   9,
		 114, 101, 115,  44,  32, 101, 114, 114,  32,  58,  61,  32,
		  82,  40, 110,  97, 109, 101,  41,  10,   9, 105, 102,  32,
		 101, 114, 114,  32,  33,  61,  32, 110, 105, 108,  32, 123,
		  10,   9,  32,  32,  32,  32, 114, 101, 116, 117, 114, 110,
		  32,  34,  34,  44,  32, 101, 114, 114,  10,   9, 125,  10,
		   9, 114, 101, 116, 117, 114, 110,  32, 115, 116, 114, 105,
		 110, 103,  40, 114, 101, 115,  41,  44,  32, 110, 105, 108,
		  10,  32,  32,  32,  32, 123, 123,  32, 101, 108, 115, 101,
		  32, 125, 125,  10,   9, 114, 101, 116, 117, 114, 110,  32,
		 115, 116, 114, 105, 110, 103,  40,  82,  40, 110,  97, 109,
		 101,  41,  41,  10,  32,  32,  32,  32, 123, 123,  32, 101,
		 110, 100,  32, 125, 125,  10, 125,  10,
	},
    
	"resources.yaml.tpl": []byte{ 
		  45,  45,  45,  10,  35,  32,  80,  97,  99, 107,  97, 103,
		 101,  32, 110,  97, 109, 101,  32, 102, 111, 114,  32, 114,
		 101, 115, 117, 108, 116,  32,  46, 103, 111,  10, 112,  97,
		  99, 107,  97, 103, 101,  58,  32, 109,  97, 105, 110,  10,
		  10,  35,  32, 105, 103, 110, 111, 114, 101,  32, 112,  97,
		 116, 116, 101, 114, 110, 115,  32, 111, 114,  32, 110,  97,
		 109, 101, 115,  10, 105, 103, 110, 111, 114, 101,  58,  10,
		  32,  45,  32,  94, 114, 101, 115, 111, 117, 114,  99, 101,
		 115,  92,  46, 121,  97, 109, 108,  36,  10,  10,  35,  32,
		  67, 111, 100, 101, 103, 101, 110, 101, 114,  97, 116, 101,
		 100,  32, 102, 117, 110,  99, 116, 105, 111, 110, 115,  32,
		 114, 101, 116, 117, 114, 110, 115,  32, 101, 114, 114, 111,
		 114,  32,  40, 111, 114,  32, 112,  97, 110, 105,  99,  41,
		  10, 117, 115, 101,  95, 101, 114, 114, 111, 114,  58,  32,
		 102,  97, 108, 115, 101,  10,  46,  46,  46,  10,
	},
    
}

// R - return resource (as blob) by name
func R(name string) ([]byte) {

	data, ok := resData[name]

	if !ok {
	    
		panic(fmt.Sprintf("Unknown resource: %s", name))
	    
	}
	
	return data
}

// Rs - return resource (as string) by name
func Rs(name string) (string) {
    
	return string(R(name))
    
}

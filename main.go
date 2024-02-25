package main

import "image-shuffler/shuffle"

var testImagePath = "testdata/test_normal.png"
var testShufflePath = "testdata/test_shuffle.png"
var testUnshufflePath = "testdata/test_unshuffle.png"

func main() {
	shuffle.Process(testImagePath, testShufflePath)
}

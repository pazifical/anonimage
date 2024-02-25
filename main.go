package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/pazifical/anonimage/shuffle"
)

var inputPath string
var outputPath string
var shuffleMode bool
var unshuffleMode bool

func init() {
	flag.StringVar(&inputPath, "input", "", "Input file path")
	flag.StringVar(&outputPath, "output", "", "Output file path")
	flag.BoolVar(&shuffleMode, "shuffle", false, "Run in shuffle mode")
	flag.BoolVar(&unshuffleMode, "unshuffle", false, "Run in unshuffle mode")
	flag.Parse()
}

func main() {
	err := validate()
	if err != nil {
		fmt.Printf("ERROR: validating command line arguments: %v\n", err)
	}

	if shuffleMode {
		err = shuffle.Process(inputPath, outputPath, "shuffle")
	} else if unshuffleMode {
		err = shuffle.Process(inputPath, outputPath, "unshuffle")
	}
	if err != nil {
		fmt.Printf("ERROR: processing: %v\n", err)
	}
}

func validate() error {
	if inputPath == "" || outputPath == "" {
		return errors.New("input and output path have to be set")
	}

	if !shuffleMode && !unshuffleMode {
		return errors.New("please provide a command line argument")
	} else if shuffleMode && unshuffleMode {
		return errors.New("please provide only one mode")
	}

	return nil
}

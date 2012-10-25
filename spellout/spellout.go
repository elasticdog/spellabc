// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/elasticdog/spellalpha"
	"github.com/gaal/go-options/options"
	"os"
)

// flags is the options specification for processing command-line flags
const flags = `
spellout -- convert characters into spelling alphabet code words

Usage: spellout [OPTIONS] STRING...
--
v,verbose  include input characters along with output
h,help     display this help message
`

func main() {
	spec := options.NewOptions(flags)
	opt := spec.Parse(os.Args[1:])
	nonFlags := append(opt.Extra, opt.Leftover...)

	switch {
	case opt.GetBool("help"):
		spec.PrintUsageAndExit("")
	case len(nonFlags) < 1:
		spec.PrintUsageAndExit("You must supply a STRING to convert.")
	}

	for _, input := range nonFlags {
		if opt.GetBool("verbose") {
			fmt.Printf("%s  =  ", input)
		}
		spellalpha.PrintPhonetic(input)
	}
}

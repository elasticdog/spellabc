// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/elasticdog/spellalpha"
	"github.com/gaal/go-options/options"
	"os"
	"strings"
)

// flags is the options specification for processing command-line flags
const flags = `
spellout -- convert characters into spelling alphabet code words

Available alphabets include JAN (Joint Army/Navy), LAPD, NATO,
WesternUnion, and UsFinancial.

Usage: spellout [OPTIONS] STRING...
--
a,alphabet=  alphabet output is encoded in [NATO]
v,verbose    include input characters along with output
h,help       display this help message
`

func main() {
	spec := options.NewOptions(flags)
	opt := spec.Parse(os.Args[1:])
	nonFlags := append(opt.Extra, opt.Leftover...)

	if opt.GetBool("help") {
		spec.PrintUsageAndExit("")
	}

	var abc spellalpha.Encoding
	switch strings.ToLower(opt.Get("alphabet")) {
	case "jan":
		abc = *spellalpha.JanAlphabet
	case "lapd":
		abc = *spellalpha.LapdAlphabet
	case "nato":
		abc = *spellalpha.NatoAlphabet
	case "westernunion":
		abc = *spellalpha.WesternUnionAlphabet
	case "usfinancial":
		abc = *spellalpha.UsFinancialAlphabet
	default:
		spec.PrintUsageAndExit("Unknown alphabet: " + opt.Get("alphabet"))
	}

	if len(nonFlags) < 1 {
		spec.PrintUsageAndExit("You must supply a STRING to convert.")
	}

	for _, input := range nonFlags {
		if opt.GetBool("verbose") {
			fmt.Printf("%s  =  ", input)
		}
		fmt.Println(abc.Encode(input))
	}
}

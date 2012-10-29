spellabc.go & spellout
======================

spellabc.go is a [GoLang](http://golang.org/) package for converting
strings into their equivalent
[spelling alphabet](https://en.wikipedia.org/wiki/Spelling_alphabet)
(a.k.a. phonetic alphabet) code word representation. Useful for oral
communication (e.g., spelling things out over the phone), the package
supports a number of common alphabets as well as the ability to use custom
encodings.

For characters classified as letters, spellabc will return
a lowercase/uppercase code word based on the capitalization. Unknown
characters are simply returned without conversion.

`spellout` is an included command-line utility (also written in Go) that
serves as an example application of the spellabc package.

Usage
=====

    Usage: spellout [OPTIONS] STRING...

    Available alphabets include JAN (Joint Army/Navy), LAPD, NATO,
    WesternUnion, and UsFinancial.

      -a, --alphabet=  alphabet output is encoded in [NATO]
      -v, --verbose    include input characters along with output
      -h, --help       display this help message

Each string will have its output printed on a separate line, and spellout
will honor using `--` to stop interpreting the subsequent arguments as
options.

Examples
========

    $ spellout Example123
    ECHO x-ray alpha mike papa lima echo One Two Three

    $ spellout --alphabet UsFinancial Example123
    EDDIE xavier adam mary peter larry eddie One Two Three

    $ spellout --verbose Aaron "Bull Schaefer"
    Aaron  =  ALPHA alpha romeo oscar november
    Bull Schaefer  =  BRAVO uniform lima lima Space SIERRA charlie hotel alpha echo foxtrot echo romeo

    $ spellout --verbose -- --help
    --help  =  Dash Dash hotel echo lima papa

    $ spellout aΦx
    alpha Φ x-ray

And if you want to pass `spellout` the output from another command or the
contents of a file, just pipe it through `xargs`:

    $ cat secrets | xargs spellout --verbose --
    4PN%mAnt  =  Four PAPA NOVEMBER Percent mike ALPHA november tango
    5Jzd}y(d  =  Five JULIET zulu delta RightBrace yankee LeftParens delta
    BTW{2J~l  =  BRAVO TANGO WHISKEY LeftBrace Two JULIET Tilde lima

One last tip; when the contents of your file contains spaces and you only
want to separate based on newlines:

    $ cat example.txt
    Line one
    Line two
    $ cat example.txt | tr '\n' '\0' | xargs -0 spellout --verbose --
    Line one  =  LIMA india november echo Space oscar november echo
    Line two  =  LIMA india november echo Space tango whiskey oscar

Build Instructions
==================

Assuming that you already have Go installed, have set a proper `$GOROOT`
environment variable, and are inside of a workspace defined in your
`$GOPATH`, you can just use Go itself to download and build the spellabc
package:

    $ go get github.com/elasticdog/spellabc

Then import the package in your source file as usual:

    import (
    	"github.com/elasticdog/spellabc"
    )

Likewise, to build the `spellout` binary:

    $ go get github.com/elasticdog/spellabc/spellout

Then feel free to move/copy the `spellout` binary file to wherever you like:

    $ cp $GOPATH/bin/spellout /usr/local/bin/spellout

Credits
=======

spellabc was inspired by the output from the no-longer-in-existence
[WinGuides Secure Password Generator](http://www.winguides.com/security/password.php)
that disappeared back in January 2007, and the similarly inspired
[Lingua::Alphabet::Phonetic::Password](http://search.cpan.org/~jfitz/Lingua-Alphabet-Phonetic-Password-0.11/lib/Lingua/Alphabet/Phonetic/Password.pm)
Perl module written by James FitzGibbon.

License
=======

spellabc is provided under the terms of the
[MIT License](http://www.opensource.org/licenses/MIT).

Copyright &copy; 2012 [Aaron Bull Schaefer](mailto:aaron@elasticdog.com)

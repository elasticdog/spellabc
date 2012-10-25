// Copyright (c) 2012 Aaron Bull Schaefer
// This source code is provided under the terms of the MIT License
// that can be be found in the LICENSE file.

/*
Spellout converts characters into their equivalent NATO spelling alphabet
code words. Useful for when you have to spell things out over the phone.

It will currently handle ASCII printable characters, and for letters, will
return a lowercase/uppercase code word based on the capitalization.
Unknown characters are simply returned without conversion.

Usage:
    spellout [OPTIONS] STRING...

The OPTIONS are:
    -v, --verbose
        include input characters along with output
    -h, --help
        display this help message

Each string will have its output printed on a separate line, and spellout
will honor using `--` to stop interpreting the subsequent arguments as
options.

Examples

    $ spellout Example123
    ECHO x-ray alpha mike papa lima echo One Two Three

    $ spellout --verbose Aaron "Bull Schaefer"
    Aaron  =  ALPHA alpha romeo oscar november
    Bull Schaefer  =  BRAVO uniform lima lima Space SIERRA charlie hotel alpha echo foxtrot echo romeo

    $ spellout --verbose -- --help
    --help  =  Dash Dash hotel echo lima papa

    $ spellout aΦx
    alpha Φ x-ray

And if you want to pass spellout the output from another command or the
contents of a file, just pipe it through xargs:

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
*/
package documentation

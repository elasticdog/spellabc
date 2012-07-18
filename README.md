go-phonetic
===========

A command-line utility (written in [GoLang](http://golang.org/)) for
converting characters into their equivalent NATO [spelling
alphabet](https://en.wikipedia.org/wiki/Spelling_alphabet) code words.
Useful for when you have to spell things out over the phone.

go-phonetic will currently handle ASCII printable characters, and for
letters, will return a lowercase/uppercase code word based on the
capitalization. Unknown characters are simply returned without conversion.

Usage
=====

    Usage: phonetic [OPTIONS] STRING...

      -v, --verbose  include input characters along with output
      -h, --help     display this help message

Each string will have its output printed on a separate line, and
go-phonetic will honor using `--` to stop interpreting the subsequent
arguments as options.

Examples
========

    $ phonetic Example123
    ECHO x-ray alpha mike papa lima echo One Two Three

    $ phonetic --verbose Aaron "Bull Schaefer"
    Aaron  =  ALPHA alpha romeo oscar november
    Bull Schaefer  =  BRAVO uniform lima lima Space SIERRA charlie hotel alpha echo foxtrot echo romeo

    $ phonetic --verbose -- --help
    --help  =  Dash Dash hotel echo lima papa

    $ phonetic aΦx
    alpha Φ x-ray

And if you want to pass `phonetic` the output from another command or the
contents of a file, just pipe it through `xargs`:

    $ cat secrets | xargs phonetic --verbose --
    4PN%mAnt  =  Four PAPA NOVEMBER Percent mike ALPHA november tango
    5Jzd}y(d  =  Five JULIET zulu delta RightBrace yankee LeftParens delta
    BTW{2J~l  =  BRAVO TANGO WHISKEY LeftBrace Two JULIET Tilde lima

One last tip; when the contents of your file contains spaces and you only
want to separate based on newlines:

    $ cat example.txt
    Line one
    Line two
    $ cat example.txt | tr '\n' '\0' | xargs -0 phonetic --verbose --
    Line one  =  LIMA india november echo Space oscar november echo
    Line two  =  LIMA india november echo Space tango whiskey oscar

Credits
=======

go-phonetic was inspired by the output from the no-longer-in-existence
[WinGuides Secure Password Generator](http://www.winguides.com/security/password.php)
that disappeared back in Jan 2007, and the similarly inspired
[Lingua::Alphabet::Phonetic::Password](http://search.cpan.org/~jfitz/Lingua-Alphabet-Phonetic-Password-0.11/lib/Lingua/Alphabet/Phonetic/Password.pm)
Perl module written by James FitzGibbon.

License
=======

go-phonetic is provided under the terms of [The MIT
License](http://www.opensource.org/licenses/MIT).

Copyright &copy; 2012, [Aaron Bull Schaefer](mailto:aaron@elasticdog.com).

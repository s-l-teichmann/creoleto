# creoleto

A little tool to export files written in the
[Creole](http://www.wikicreole.org/wiki/HowToUseThisWiki) markup language
to different formats.  
Currently planned supported outputs are XHTML and LaTex.

**Attention: This is work in progress and not ready to use!**

Limitations (due to limits in the grammar):
* Nested lists are not parsed.
* Links are not parsed.
* Only **bold** and *italics* emphases are supported.

## Build

You need a working [Go](https://golang.org) environment.
Tested successfully with Go 1.9.

```bash
go get github.com/s-l-teichmann/creoleto
```

Put the resulting `creoleto` binary into your `PATH`.

## Usage

```bash
creoleto -standalone -format latex < data/text.creole > converted.tex
```

## Creole grammar

The GPLed Creole 1.0 ANTLR3 grammar was taken from
[here](http://wikicreole.cvs.sourceforge.net/viewvc/wikicreole/XmlCreator/creole10.g)
and was modified to be usable with ANTLR4.  
You need a pretty recent [ANTLR4](http://www.antlr.org/download.html)
(tested successfully with 4.6+) to create the Go parser code:

```bash
antlr4 -Dlanguage=Go -o parser Creole10.g4
```

To build the resulting code you need Go ANTLR4 runtime:

```bash
go get github.com/antlr/antlr4/runtime/Go/antlr
```

## License

This is Free Software covered by the terms of the [MIT license](LICENSE).  
Copyright (c) 2017 by Intevation GmbH

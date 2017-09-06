# creoleto

A little tool to export files written in the Creole markup language
to different formats.  
Currently planned supported outputs are XHTML and LaTex.

## Build

You need a working [Go](https://golang.org) environment.
Tested successfully with Go 1.9.

    $ go get github.com/s-l-teichmann/creoleto

Put the resulting `creoleto` binary into your `PATH`

## Usage

    $ creoleto < data/wikicreole1.creole > converted.tex

## License

This is Free Software covered by the terms of the MIT license.  
(c) 2017 by Intevation GmbH

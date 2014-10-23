# Goampl: [Golang](https://golang.org/) interface to [AMPL](http://ampl.com/)

## Installation (Linux)

Requires [Linuxbrew](https://github.com/Homebrew/linuxbrew) for easy installation of ASL:

	brew tap homebrew/science

Must change line 28 of asl.rb from 
	
	libtool_cmd = ["ld", "-shared"]
to
	
	libtool_cmd = ["gcc", "-shared"]

Then:

	brew install asl

In goampl.go package alter CFLAGS and LDFLAGS accordingly. Then it should be able to run.
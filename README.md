<!---
[![GoDoc](https://godoc.org/github.com/logavanc/largs?status.svg)](https://godoc.org/github.com/logavanc/largs)
[![Build Status](https://travis-ci.org/logavanc/largs.svg?branch=master)](https://travis-ci.org/logavanc/largs)
[![Coverage Status](https://img.shields.io/coveralls/logavanc/largs.svg)](https://coveralls.io/r/logavanc/largs)
-->
[![Go Report Card](https://goreportcard.com/badge/github.com/logavanc/shortpath)](https://goreportcard.com/report/github.com/logavanc/shortpath)


The "shortpath" tool...
======================

...is a command line utility (written in [Go](http://golang.org)) that returns a string representing the current working directory where the name of each parent directory has been shortened to the smallest uniquely identifiable string for the directory in which it resides. The primary intended use case for this utility is to construct the current working directory in the command line prompt.  For example, a normal prompt would contain the full path to the current working directory (pwd) in the prompt, but with my [`shortpath`](https://github.com/logavanc/shortpath) utility, the prompt is shortened considerably without removing so much information that confusion could occur.

![The "shortpath" utility in use.](/images/example.png)


At the moment, it simply walks up the current working directory path and finds the shortest string for each directory that still represents that directory uniquely.  This operation has the possibility of being expensive, and since it happens every time you run the command (every time the prompt is shown), it really needs optimized somehow...

Installing
----------

    go get github.com/logavanc/shortpath

Documentation
-------------

See [documentation on godoc.org](https://godoc.org/github.com/logavanc/shortpath).

License
-------

GNU GENERAL PUBLIC LICENSE Version 3.

See the [LICENSE](LICENSE) file for details and
[this site](https://www.gnu.org/licenses/rms-why-gplv3.html) for reasoning.

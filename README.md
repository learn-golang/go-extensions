Go-extensions
=============

This repository reflects attempts to use Golang functions inside Python through C-extensions and shared libs.

HOW-TO
------

In order to build code under `alternative` directory you need:

 * setup GOPATH (export GOPATH=`pwd`)
 * build code (go build -buildmode=c-shared -o libadd.so libadd.go)

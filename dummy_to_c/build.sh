#!/usr/bin/env bash

go build -buildmode=c-shared -o dummy.so dummy.go
gcc -Wall -Wformat -o main main.c ./dummy.so
./main
rm -fr dummy.so dummy.h main
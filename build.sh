#!/bin/bash

export GOPROXY=https://gocenter.io

go mod download
go build png2escpos.go
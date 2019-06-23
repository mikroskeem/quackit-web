#!/bin/sh
GOOS=js
GOARCH=wasm
export GOOS
export GOARCH

"${@}"

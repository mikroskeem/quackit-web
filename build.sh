#!/bin/sh
if [ ! -f "wasm_exec.js" ]; then
    cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
fi

./goenv.sh go build -o main.wasm

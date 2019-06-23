#!/bin/sh
set -e
if [ ! -f "wasm_exec.js" ]; then
    cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
fi

./goenv.sh go build -o main.wasm

tar -cf app.tar \
    index.html \
    main.wasm \
    wasm_exec.js

// +build js,wasm

package main

import (
	"fmt"
	"sort"
	"strings"
	"syscall/js"

	"github.com/mikroskeem/quackit"
)

var done chan struct{}

func main() {
	q := new(quackit.Quackit)

	// Set up config parser function
	parseConfigCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		config := getConfigText()

		if err := q.ParseString(config); err != nil {
			fmt.Printf("Encountered an error: %s\n", err)
		}
		fmt.Printf("Parsed commands: %d\n", len(q.ParsedCommands()))

		displayParsed(q.ParsedCommands())

		return nil
	})
	js.Global().Get("document").Call("querySelector", "button#configsubmit").Call("addEventListener", "click", parseConfigCb)

	fmt.Println("Initialized!")

	// Keeps Go program running forever
	<-done
}

func displayParsed(parsedTokens [][]quackit.Token) {
	pre := js.Global().Get("document").Call("querySelector", "pre#parsed")
	minify := getMinify()

	var generated strings.Builder
	packedCommands := []string{}

	for _, commandTokens := range parsedTokens {
		var sb strings.Builder
		for _, token := range commandTokens {
			switch token.GetType() {
			case quackit.TokenTypeWord:
				sb.WriteString(token.(quackit.WordToken).Word)
			case quackit.TokenTypeString:
				sb.WriteString(fmt.Sprintf(`"%s"`, token.(quackit.StringToken).Value))
			}
			sb.WriteString(" ")
		}
		packedCommands = append(packedCommands, strings.TrimSpace(sb.String()))
	}

	sort.Strings(packedCommands)
	for _, command := range packedCommands {
		if minify {
			generated.WriteString(fmt.Sprintf("%s;", command))
		} else {
			generated.WriteString(fmt.Sprintf("%s\n", command))
		}
	}

	pre.Set("innerHTML", generated.String())
}

func getConfigText() string {
	return js.Global().Get("document").Call("querySelector", "textarea#config").Get("value").String()
}

func getMinify() bool {
	return js.Global().Get("document").Call("querySelector", "input#minify").Get("checked").Bool()
}

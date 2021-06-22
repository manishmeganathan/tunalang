package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/manishmeganathan/tuna/repl"
)

const version = "v0.5.0"

func main() {
	fmt.Println(repl.TUNA2)
	fmt.Printf("The Tuna Programming Language %s [%s-%s].\n", version, strings.Title(runtime.GOOS), strings.ToUpper(runtime.GOARCH))
	fmt.Println("Welcome to the Tuna REPL. Visit www.github.com/manishmeganathan/tuna for more information.")
	repl.StartREPL(os.Stdin, os.Stdout)
}

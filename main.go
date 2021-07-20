package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/manishmeganathan/tunalang/repl"
)

const version = "v1.0.0"

func main() {
	fmt.Println(repl.TUNA2)
	fmt.Printf("The Tuna Programming Language %s [%s-%s].\n", version, strings.Title(runtime.GOOS), strings.ToUpper(runtime.GOARCH))
	fmt.Println("Welcome to the Tuna REPL. Visit www.github.com/manishmeganathan/tunalang for more information.")
	repl.StartREPL(os.Stdin, os.Stdout)
}

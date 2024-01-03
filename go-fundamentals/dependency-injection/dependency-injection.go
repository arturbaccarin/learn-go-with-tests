package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func PrintToConsole() {
	Greet(os.Stdout, "Elodie")
}

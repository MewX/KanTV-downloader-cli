package kantv

import (
	"fmt"
	"testing"
)

func BenchmarkCli(b *testing.B) {
	fmt.Println("Test placeholder starts!")
	Cli()
	fmt.Println("Test placeholder finished!")
}

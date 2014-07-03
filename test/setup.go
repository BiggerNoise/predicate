// use test.sh

package main

import (
	"os"
	"strings"

	_ "github.com/biggernoise/predicate"
	"github.com/clipperhouse/gen/typewriter"
)

func main() {
	// don't let bad test or gen'd files get us stuck
	filter := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go") && !strings.HasSuffix(f.Name(), "_predicate.go")
	}

	a, err := typewriter.NewAppFiltered("+test", filter)
	if err != nil {
		panic(err.Error())
	}
	a.WriteAll()
}

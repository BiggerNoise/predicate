package predicate

import (
	"fmt"
	"io"

	"github.com/clipperhouse/gen/typewriter"
)

func init() {
	err := typewriter.Register(NewPredicateWriter())
	if err != nil {
		panic(err)
	}
}

type PredicateWriter struct {
	// since typewriter.Type is not comparable, key by typewriter.Type.String()
	tagsByType map[string]typewriter.Tag
}

func NewPredicateWriter() *PredicateWriter {
	return &PredicateWriter{
		tagsByType: make(map[string]typewriter.Tag),
	}
}

func (c *PredicateWriter) Name() string {
	return "predicate"
}

// Validates that the tag on the gen type has correctly instructed this
// typewriter to generate code.  If the first return is false, then
// none of the write methods are called.
func (c *PredicateWriter) Validate(twt typewriter.Type) (bool, error) {
	tag, found, err := twt.Tags.ByName("predicates")

	if !found || err != nil {
		return false, err
	}

	// must include at least one item that we recognize
	any := false
	for _, item := range tag.Items {
		if templates.Contains(item) {
			// found one, move on
			any = true
			break
		}
	}

	if !any {
		// not an error, but irrelevant
		return false, nil
	}

	c.tagsByType[twt.String()] = tag
	return true, nil
}

// Write headers for the generated file.  This would include licencing, credits
// or anythign else that you wanted to put in there.  This can depend upon the
// items selected.
func (c *PredicateWriter) WriteHeader(wrt io.Writer, twt typewriter.Type) {
	s := `// Generated Predicates 
	`

	wrt.Write([]byte(s))
	s = `// See http://clipperhouse.github.io/gen for documentation on gen

`
	wrt.Write([]byte(s))
}

// Return the imports for the generated file.  The gen library handles writing the actual
// output.
func (c *PredicateWriter) Imports(twt typewriter.Type) (result []typewriter.ImportSpec) {
	// none
	return result
}

// Write the actual body.
func (c *PredicateWriter) WriteBody(wrt io.Writer, twt typewriter.Type) {
	tag := c.tagsByType[twt.String()] // validated above

	// Always include the predicate template.  We wouldn't be here if there wasn't at
	// least one valid item being written
	items := append([]string{"predicate"}, tag.Items...)

	for _, item := range items {
		tmpl, err := templates.Get(item) // validate above to avoid err check here?
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = tmpl.Execute(wrt, twt)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return
}

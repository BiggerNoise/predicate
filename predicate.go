package predicate

import "github.com/clipperhouse/gen/typewriter"

func init() {
	err := typewriter.Register(NewPredicateWriter())
	if err != nil {
		panic(err)
	}
}

func NewPredicateWriter() *SimpleTypewriter {
	return NewSimpleTypewriter("predicate").
		WithTemplates(predicateTemplates).
		WithHeader("// Simple predicates with boolean combinations").
		WithHeader("// See http://clipperhouse.github.io/gen for documentation on gen").
		WithItemHeader("Not", `
		// Not() is syntacticly a bit ugly.  Ideally, you would say something 
		// like Not(other). Because there's no overloading, you have to settle
		// for other.Not()`)
}

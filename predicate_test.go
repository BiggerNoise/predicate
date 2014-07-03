package predicate

import (
	"testing"

	"github.com/clipperhouse/gen/typewriter"
)

func TestValidate(t *testing.T) {
	g := NewPredicateWriter()

	pkg := typewriter.NewPackage("dummy", "SomePackage")

	typ := typewriter.Type{
		Package: pkg,
		Name:    "SomeType",
		Tags:    typewriter.Tags{},
	}

	write, err := g.Validate(typ)

	if write {
		t.Errorf("no 'predicates' tag should not write")
	}

	if err != nil {
		t.Error(err)
	}

	typ2 := typewriter.Type{
		Package: pkg,
		Name:    "SomeType2",
		Tags: typewriter.Tags{
			typewriter.Tag{
				Name:  "predicates",
				Items: []string{},
			},
		},
	}

	write2, err2 := g.Validate(typ2)

	if write2 {
		t.Errorf("empty 'predicates' tag should not write")
	}

	if err2 != nil {
		t.Error(err)
	}

	typ3 := typewriter.Type{
		Package: pkg,
		Name:    "SomeType3",
		Tags: typewriter.Tags{
			typewriter.Tag{
				Name:  "predicates",
				Items: []string{"And", "Foo"},
			},
		},
	}

	write3, err3 := g.Validate(typ3)

	if !write3 {
		t.Errorf("'predicates' tag with And should write (and ignore others)")
	}

	if err3 != nil {
		t.Error(err)
	}
}

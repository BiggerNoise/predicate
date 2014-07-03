package predicate

import (
	"github.com/clipperhouse/gen/typewriter"
)

var predicateTemplates = typewriter.TemplateSet{
	Common: &typewriter.Template{
		Text: `// {{.Name}}Predicate is a function that accepts a {{.Pointer}}{{.Name}} and returns a bool.  Use this type where you would use func({{.Pointer}}{{.Name}}) bool.
type {{.Name}}Predicate func(item {{.Pointer}}{{.Name}}) bool
`},
	"And": &typewriter.Template{
		Text: `// And combines two predicates into a new predicate that is satisfied if both of the original predicates are satisfied
func (rcv {{.Name}}Predicate) And(other {{.Name}}Predicate) {{.Name}}Predicate {
	return func (item {{.Pointer}}{{.Name}}) bool {
		return rcv(item) && other(item)
	}
}
`},
	"Or": &typewriter.Template{
		Text: `
// Or combines two predicates into a new predicate that is satisfied if either of the original predicates is satisfied
func (rcv {{.Name}}Predicate) Or(other {{.Name}}Predicate) {{.Name}}Predicate {
	return func (item {{.Pointer}}{{.Name}}) bool {
		return rcv(item)|| other(item)
	}
}
`},
	"Not": &typewriter.Template{
		Text: `
// Not inverts a predicate that is satisfied if the original predicates is not satisfied
func (rcv {{.Name}}Predicate) Not() {{.Name}}Predicate {
	return func (item {{.Pointer}}{{.Name}}) bool {
		return !rcv(item)
	}
}
`},
}

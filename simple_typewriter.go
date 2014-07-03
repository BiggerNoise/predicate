package predicate

import (
	"fmt"
	"io"
	"strings"

	"github.com/clipperhouse/gen/typewriter"
	"github.com/clipperhouse/inflect"
)

const Common = "Common"

// A simple typewriter is intended to be a very simple bridge between a template set
// and the typewriter interface expected by clipperhouse/gen.  It is appropriate iff
// you have a 1:1 correspondence between the names of the templates and the items you
// wish to generate.
type SimpleTypewriter struct {
	name               string
	itemsForTypeName   map[string][]string
	templateSet        typewriter.TemplateSet
	headersForItemType map[string][]string
}

func NewSimpleTypewriter(name string) *SimpleTypewriter {
	return &SimpleTypewriter{
		name:               name,
		itemsForTypeName:   make(map[string][]string),
		headersForItemType: make(map[string][]string),
	}
}

// --------------------- Configuration ---------------------

func (tw *SimpleTypewriter) WithTemplates(templates typewriter.TemplateSet) *SimpleTypewriter {
	tw.templateSet = templates
	return tw
}

func (tw *SimpleTypewriter) WithHeader(headerText string) *SimpleTypewriter {
	return tw.WithItemHeader(Common, headerText)
}

func (tw *SimpleTypewriter) WithItemHeader(itemTag string, headerText string) *SimpleTypewriter {
	tw.headersForItemType[itemTag] = append(tw.headersForItemType[itemTag], headerText)
	return tw
}

// --------------------- Implementation ---------------------
func plural(input string) (result string) {
	result = inflect.Pluralize(input)
	if result == input {
		result += "s"
	}
	return
}

func (tw *SimpleTypewriter) itemsToGenerate(names []string) []string {
	items := make([]string, 0, len(names)+1)
	for _, name := range names {
		if tw.templateSet.Contains(name) {
			items = append(items, name)
		}
	}
	if len(items) > 0 && tw.templateSet.Contains(Common) {
		return append([]string{Common}, items...)
	} else {
		return items
	}
}

func (tw *SimpleTypewriter) Name() string {
	return tw.name
}

// Validates that the tag on the gen type has correctly instructed this
// typewriter to generate code.  If the first return value is false, then
// none of the write methods are called.
func (tw *SimpleTypewriter) Validate(twt typewriter.Type) (bool, error) {
	tag, found, err := twt.Tags.ByName(plural(tw.name))

	if !found || err != nil {
		return false, err
	}

	items := tw.itemsToGenerate(tag.Items)
	fmt.Printf("%s: items for %s: %v\n", tw.Name(), twt.String(), items)

	if len(items) == 0 {
		// not an error, but irrelevant to generation
		return false, nil
	}

	tw.itemsForTypeName[twt.String()] = items
	return true, nil
}

// Write headersForItemType for the generated file.  This would include licensing, credits
// or anything else that you wanted to put in there.  This can depend upon the
// items selected.
func (tw *SimpleTypewriter) WriteHeader(wrt io.Writer, twt typewriter.Type) {
	for _, item := range tw.itemsForTypeName[twt.String()] {
		tw.writeHeaderForItem(item, wrt)
	}
}

func (tw *SimpleTypewriter) writeHeaderForItem(itemName string, wrt io.Writer) {
	headers, present := tw.headersForItemType[itemName]
	if !present {
		return
	}
	wrt.Write([]byte(strings.Join(headers, "\n")))
}

// Return the imports for the generated file.  The gen library handles writing the actual
// output.
func (tw *SimpleTypewriter) Imports(twt typewriter.Type) (result []typewriter.ImportSpec) {
	// none, not handled just yet
	return result
}

// Write the actual body.
func (tw *SimpleTypewriter) WriteBody(wrt io.Writer, twt typewriter.Type) {
	for _, item := range tw.itemsForTypeName[twt.String()] {
		tw.writeTemplateBody(item, wrt, twt)
	}
	return
}

func (tw *SimpleTypewriter) writeTemplateBody(templateName string, wrt io.Writer, twt typewriter.Type) {
	template, err := tw.templateSet.Get(templateName)
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(wrt, twt)
	if err != nil {
		fmt.Println(err)
	}
}

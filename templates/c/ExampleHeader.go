/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package templates

// ExampleHeaderTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func ExampleHeaderTemplate() string {
	var tmpl = "#ifndef {{ .HeaderName }}\n" +
		"#define {{ .HeaderName }}\n" +
		"\n" +
		"\n" +
		"#endif"
	return tmpl
}

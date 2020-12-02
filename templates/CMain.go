/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package templates

// CMainTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func CMainTemplate() string {
	var tmpl = "#include <stdio.h>\n" +
		"\n" +
		"int main() {\n" +
		"\tprintf(\"Hello World\\n\");\n" +
		"\treturn 0;\n" +
		"}"
	return tmpl
}

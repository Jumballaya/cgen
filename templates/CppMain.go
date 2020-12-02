/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package templates

// CppMainTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func CppMainTemplate() string {
	var tmpl = "#include <iostream>\n" +
		"#include \"../inc/Vector.h\"\n" +
		"\n" +
		"\n" +
		"int main() {\n" +
		"\n" +
		"  Vector<int> v = 10;\n" +
		"\n" +
		"  for (int i = 0; i < v.size(); i++) {\n" +
		"    v[i] = i;\n" +
		"  }\n" +
		"\n" +
		"  for (int i = 0; i < v.size(); i++) {\n" +
		"    std::cout << v[i] << \"\\n\";\n" +
		"  }\n" +
		"\n" +
		"  return(0);\n" +
		"}"
	return tmpl
}

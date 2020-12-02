/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package templates

// CppVectorHeaderTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func CppVectorHeaderTemplate() string {
	var tmpl = "#ifndef H_VECTOR\n" +
		"#define H_VECTOR\n" +
		"\n" +
		"template<typename T>\n" +
		"class Vector {\n" +
		"    private:\n" +
		"        T* elem;\n" +
		"        int _size;\n" +
		"\n" +
		"    public:\n" +
		"        Vector(int s);\n" +
		"        ~Vector() { delete[] elem; }\n" +
		"\n" +
		"        Vector(const Vector& v);\n" +
		"        Vector& operator=(const Vector& v);\n" +
		"\n" +
		"        Vector(Vector&& v);\n" +
		"        Vector& operator=(Vector&& v);\n" +
		"\n" +
		"        T& operator[](int i);\n" +
		"        const T& operator[](int i) const;\n" +
		"\n" +
		"        T* begin(Vector<T>& v);\n" +
		"        T* end(Vector<T>& v);\n" +
		"\n" +
		"        int size() const { return _size; }\n" +
		"};\n" +
		"\n" +
		"template<typename T>\n" +
		"Vector<T>::Vector(const Vector<T>& v) : elem{new T[v._size]}, _size{v._size} {\n" +
		"    for (int i = 0; i < _size; i++) {\n" +
		"        elem[i] = v.elem[i];\n" +
		"    }\n" +
		"}\n" +
		"\n" +
		"template<typename T>\n" +
		"Vector<T>& Vector<T>::operator=(const Vector<T>& v) {\n" +
		"    T* p = new T[v._size];\n" +
		"    for (int i = 0; i < v._size; i++) {\n" +
		"        p[i] = v.elem[i];\n" +
		"    }\n" +
		"    delete[] elem;\n" +
		"    elem = p;\n" +
		"    _size = v._size;\n" +
		"    return *this;\n" +
		"}\n" +
		"\n" +
		"template<typename T>\n" +
		"Vector<T>::Vector(int s) {\n" +
		"    if (s < 0) {\n" +
		"        throw std::out_of_range{\"Can't initialize Vector of negative size\"};\n" +
		"    }\n" +
		"    elem = new T[s];\n" +
		"    _size = s;\n" +
		"}\n" +
		"\n" +
		"\n" +
		"template<typename T>\n" +
		"Vector<T>::Vector(Vector&& v) : elem{v.elem}, _size{v._size} {\n" +
		"    v.elem = nullptr;\n" +
		"    v._size = 0;\n" +
		"}\n" +
		"\n" +
		"template<typename T>\n" +
		"const T& Vector<T>::operator[](int i) const {\n" +
		"    if (i < 0 || size() <= i) {\n" +
		"        throw std::out_of_range{\"Vector::operator[]\"};\n" +
		"    }\n" +
		"    return elem[i];\n" +
		"}\n" +
		"\n" +
		"template<typename T>\n" +
		"T& Vector<T>::operator[](int i) {\n" +
		"    if (i < 0 || size() <= i) {\n" +
		"        throw std::out_of_range{\"Vector::operator[]\"};\n" +
		"    }\n" +
		"    return elem[i];\n" +
		"}\n" +
		"\n" +
		"template<typename T>\n" +
		"T* begin(Vector<T>& v) {\n" +
		"    return v.size() ? &v[0] : nullptr;\n" +
		"}\n" +
		"\n" +
		"template<typename T>\n" +
		"T* end(Vector<T>& v) {\n" +
		"    return begin(v) + v.size();\n" +
		"}\n" +
		"\n" +
		"\n" +
		"#endif"
	return tmpl
}

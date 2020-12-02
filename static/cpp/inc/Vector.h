#ifndef H_VECTOR
#define H_VECTOR

template<typename T>
class Vector {
    private:
        T* elem;
        int _size;

    public:
        Vector(int s);
        ~Vector() { delete[] elem; }

        Vector(const Vector& v);
        Vector& operator=(const Vector& v);

        Vector(Vector&& v);
        Vector& operator=(Vector&& v);

        T& operator[](int i);
        const T& operator[](int i) const;

        T* begin(Vector<T>& v);
        T* end(Vector<T>& v);

        int size() const { return _size; }
};

template<typename T>
Vector<T>::Vector(const Vector<T>& v) : elem{new T[v._size]}, _size{v._size} {
    for (int i = 0; i < _size; i++) {
        elem[i] = v.elem[i];
    }
}

template<typename T>
Vector<T>& Vector<T>::operator=(const Vector<T>& v) {
    T* p = new T[v._size];
    for (int i = 0; i < v._size; i++) {
        p[i] = v.elem[i];
    }
    delete[] elem;
    elem = p;
    _size = v._size;
    return *this;
}

template<typename T>
Vector<T>::Vector(int s) {
    if (s < 0) {
        throw std::out_of_range{"Can't initialize Vector of negative size"};
    }
    elem = new T[s];
    _size = s;
}


template<typename T>
Vector<T>::Vector(Vector&& v) : elem{v.elem}, _size{v._size} {
    v.elem = nullptr;
    v._size = 0;
}

template<typename T>
const T& Vector<T>::operator[](int i) const {
    if (i < 0 || size() <= i) {
        throw std::out_of_range{"Vector::operator[]"};
    }
    return elem[i];
}

template<typename T>
T& Vector<T>::operator[](int i) {
    if (i < 0 || size() <= i) {
        throw std::out_of_range{"Vector::operator[]"};
    }
    return elem[i];
}

template<typename T>
T* begin(Vector<T>& v) {
    return v.size() ? &v[0] : nullptr;
}

template<typename T>
T* end(Vector<T>& v) {
    return begin(v) + v.size();
}


#endif
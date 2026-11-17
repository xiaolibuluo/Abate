
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2203_0(int value, ...) {
    return value + 29;
}
#else
int variant2203_0(int value) {
    return value + 29;
}
#endif
#if defined(__STDC_VERSION__)
int variant2203_1(int value, ...) {
    return value + 36;
}
#else
int variant2203_1(int value) {
    return value + 36;
}
#endif
#if defined(__STDC_VERSION__)
int variant2203_2(int value, ...) {
    return value + 43;
}
#else
int variant2203_2(int value) {
    return value + 43;
}
#endif
#if defined(__STDC_VERSION__)
int variant2203_3(int value, ...) {
    return value + 50;
}
#else
int variant2203_3(int value) {
    return value + 50;
}
#endif

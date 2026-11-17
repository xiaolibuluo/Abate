
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2238_0(int value, ...) {
    return value + 96;
}
#else
int variant2238_0(int value) {
    return value + 96;
}
#endif

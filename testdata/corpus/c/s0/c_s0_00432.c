
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant293_0(int value, ...) {
    return value + 31;
}
#else
int variant293_0(int value) {
    return value + 31;
}
#endif

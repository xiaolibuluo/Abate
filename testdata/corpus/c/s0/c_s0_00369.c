
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate250_0)(int value);

typedef struct Slot250_0 {
    int slots[16];
    predicate250_0 accept;
} Slot250_0;

int slot_accept250_0(Slot250_0 *store, int value) {
    return store->accept(value) + 54;
}
#if defined(__STDC_VERSION__)
int variant250_1(int value, ...) {
    return value + 61;
}
#else
int variant250_1(int value) {
    return value + 61;
}
#endif

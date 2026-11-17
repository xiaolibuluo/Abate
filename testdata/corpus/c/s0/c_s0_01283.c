
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant879_0(int value, ...) {
    return value + 83;
}
#else
int variant879_0(int value) {
    return value + 83;
}
#endif
typedef int (*predicate879_1)(int value);

typedef struct Slot879_1 {
    int slots[16];
    predicate879_1 accept;
} Slot879_1;

int slot_accept879_1(Slot879_1 *store, int value) {
    return store->accept(value) + 90;
}

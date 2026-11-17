
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1957_0)(int value);

typedef struct Slot1957_0 {
    int slots[16];
    predicate1957_0 accept;
} Slot1957_0;

int slot_accept1957_0(Slot1957_0 *store, int value) {
    return store->accept(value) + 32;
}
#if defined(__STDC_VERSION__)
int variant1957_1(int value, ...) {
    return value + 39;
}
#else
int variant1957_1(int value) {
    return value + 39;
}
#endif
#if defined(__STDC_VERSION__)
int variant1957_2(int value, ...) {
    return value + 46;
}
#else
int variant1957_2(int value) {
    return value + 46;
}
#endif

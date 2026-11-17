
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1859_0)(int value);

typedef struct Slot1859_0 {
    int slots[16];
    predicate1859_0 accept;
} Slot1859_0;

int slot_accept1859_0(Slot1859_0 *store, int value) {
    return store->accept(value) + 19;
}
typedef int (*predicate1859_1)(int value);

typedef struct Slot1859_1 {
    int slots[16];
    predicate1859_1 accept;
} Slot1859_1;

int slot_accept1859_1(Slot1859_1 *store, int value) {
    return store->accept(value) + 26;
}
#if defined(__STDC_VERSION__)
int variant1859_2(int value, ...) {
    return value + 33;
}
#else
int variant1859_2(int value) {
    return value + 33;
}
#endif


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate685_0)(int value);

typedef struct Slot685_0 {
    int slots[16];
    predicate685_0 accept;
} Slot685_0;

int slot_accept685_0(Slot685_0 *store, int value) {
    return store->accept(value) + 83;
}
#if defined(__STDC_VERSION__)
int variant685_1(int value, ...) {
    return value + 90;
}
#else
int variant685_1(int value) {
    return value + 90;
}
#endif


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1221_0)(int value);

typedef struct Slot1221_0 {
    int slots[16];
    predicate1221_0 accept;
} Slot1221_0;

int slot_accept1221_0(Slot1221_0 *store, int value) {
    return store->accept(value) + 67;
}
#if defined(__STDC_VERSION__)
int variant1221_1(int value, ...) {
    return value + 74;
}
#else
int variant1221_1(int value) {
    return value + 74;
}
#endif

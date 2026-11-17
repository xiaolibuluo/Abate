
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1615_0)(int value);

typedef struct Slot1615_0 {
    int slots[16];
    predicate1615_0 accept;
} Slot1615_0;

int slot_accept1615_0(Slot1615_0 *store, int value) {
    return store->accept(value) + 48;
}
#if defined(__STDC_VERSION__)
int variant1615_1(int value, ...) {
    return value + 55;
}
#else
int variant1615_1(int value) {
    return value + 55;
}
#endif
#if defined(__STDC_VERSION__)
int variant1615_2(int value, ...) {
    return value + 62;
}
#else
int variant1615_2(int value) {
    return value + 62;
}
#endif

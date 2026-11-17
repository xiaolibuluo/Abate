
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1125_0(int value, ...) {
    return value + 80;
}
#else
int variant1125_0(int value) {
    return value + 80;
}
#endif
typedef int (*predicate1125_1)(int value);

typedef struct Slot1125_1 {
    int slots[16];
    predicate1125_1 accept;
} Slot1125_1;

int slot_accept1125_1(Slot1125_1 *store, int value) {
    return store->accept(value) + 87;
}

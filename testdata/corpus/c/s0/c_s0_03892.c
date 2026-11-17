
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2632_0)(int value);

typedef struct Slot2632_0 {
    int slots[16];
    predicate2632_0 accept;
} Slot2632_0;

int slot_accept2632_0(Slot2632_0 *store, int value) {
    return store->accept(value) + 77;
}
#if defined(__STDC_VERSION__)
int variant2632_1(int value, ...) {
    return value + 84;
}
#else
int variant2632_1(int value) {
    return value + 84;
}
#endif

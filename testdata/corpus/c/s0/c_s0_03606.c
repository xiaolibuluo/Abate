
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2435_0(int value, ...) {
    return value + 38;
}
#else
int variant2435_0(int value) {
    return value + 38;
}
#endif
typedef int (*predicate2435_1)(int value);

typedef struct Slot2435_1 {
    int slots[16];
    predicate2435_1 accept;
} Slot2435_1;

int slot_accept2435_1(Slot2435_1 *store, int value) {
    return store->accept(value) + 45;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1270_0(int value, ...) {
    return value + 25;
}
#else
int variant1270_0(int value) {
    return value + 25;
}
#endif
typedef int (*predicate1270_1)(int value);

typedef struct Slot1270_1 {
    int slots[16];
    predicate1270_1 accept;
} Slot1270_1;

int slot_accept1270_1(Slot1270_1 *store, int value) {
    return store->accept(value) + 32;
}

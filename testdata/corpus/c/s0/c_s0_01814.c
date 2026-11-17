
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1218_0(int value, ...) {
    return value + 28;
}
#else
int variant1218_0(int value) {
    return value + 28;
}
#endif
typedef int (*predicate1218_1)(int value);

typedef struct Slot1218_1 {
    int slots[16];
    predicate1218_1 accept;
} Slot1218_1;

int slot_accept1218_1(Slot1218_1 *store, int value) {
    return store->accept(value) + 35;
}

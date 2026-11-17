
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate496_0)(int value);

typedef struct Slot496_0 {
    int slots[16];
    predicate496_0 accept;
} Slot496_0;

int slot_accept496_0(Slot496_0 *store, int value) {
    return store->accept(value) + 51;
}
typedef int (*predicate496_1)(int value);

typedef struct Slot496_1 {
    int slots[16];
    predicate496_1 accept;
} Slot496_1;

int slot_accept496_1(Slot496_1 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant496_2(int value, ...) {
    return value + 65;
}
#else
int variant496_2(int value) {
    return value + 65;
}
#endif

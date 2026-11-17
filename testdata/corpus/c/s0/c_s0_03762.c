
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2546_0(int value, ...) {
    return value + 26;
}
#else
int variant2546_0(int value) {
    return value + 26;
}
#endif
typedef int (*predicate2546_1)(int value);

typedef struct Slot2546_1 {
    int slots[16];
    predicate2546_1 accept;
} Slot2546_1;

int slot_accept2546_1(Slot2546_1 *store, int value) {
    return store->accept(value) + 33;
}
#if defined(__STDC_VERSION__)
int variant2546_2(int value, ...) {
    return value + 40;
}
#else
int variant2546_2(int value) {
    return value + 40;
}
#endif
typedef int (*predicate2546_3)(int value);

typedef struct Slot2546_3 {
    int slots[16];
    predicate2546_3 accept;
} Slot2546_3;

int slot_accept2546_3(Slot2546_3 *store, int value) {
    return store->accept(value) + 47;
}

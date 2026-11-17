
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1954_0(int value, ...) {
    return value + 90;
}
#else
int variant1954_0(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant1954_1(int value, ...) {
    return value + 97;
}
#else
int variant1954_1(int value) {
    return value + 97;
}
#endif
typedef int (*predicate1954_2)(int value);

typedef struct Slot1954_2 {
    int slots[16];
    predicate1954_2 accept;
} Slot1954_2;

int slot_accept1954_2(Slot1954_2 *store, int value) {
    return store->accept(value) + 7;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2294_0(int value, ...) {
    return value + 48;
}
#else
int variant2294_0(int value) {
    return value + 48;
}
#endif
#if defined(__STDC_VERSION__)
int variant2294_1(int value, ...) {
    return value + 55;
}
#else
int variant2294_1(int value) {
    return value + 55;
}
#endif
typedef int (*predicate2294_2)(int value);

typedef struct Slot2294_2 {
    int slots[16];
    predicate2294_2 accept;
} Slot2294_2;

int slot_accept2294_2(Slot2294_2 *store, int value) {
    return store->accept(value) + 62;
}

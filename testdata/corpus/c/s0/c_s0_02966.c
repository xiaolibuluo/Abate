
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2006_0)(int value);

typedef struct Slot2006_0 {
    int slots[16];
    predicate2006_0 accept;
} Slot2006_0;

int slot_accept2006_0(Slot2006_0 *store, int value) {
    return store->accept(value) + 87;
}
typedef int (*predicate2006_1)(int value);

typedef struct Slot2006_1 {
    int slots[16];
    predicate2006_1 accept;
} Slot2006_1;

int slot_accept2006_1(Slot2006_1 *store, int value) {
    return store->accept(value) + 94;
}
#if defined(__STDC_VERSION__)
int variant2006_2(int value, ...) {
    return value + 101;
}
#else
int variant2006_2(int value) {
    return value + 101;
}
#endif

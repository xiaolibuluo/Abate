
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2639_0(int value, ...) {
    return value + 71;
}
#else
int variant2639_0(int value) {
    return value + 71;
}
#endif
typedef int (*predicate2639_1)(int value);

typedef struct Slot2639_1 {
    int slots[16];
    predicate2639_1 accept;
} Slot2639_1;

int slot_accept2639_1(Slot2639_1 *store, int value) {
    return store->accept(value) + 78;
}
#define MASK2639_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2639_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2639_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

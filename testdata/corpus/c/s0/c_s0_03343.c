
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2256_0(int value, ...) {
    return value + 39;
}
#else
int variant2256_0(int value) {
    return value + 39;
}
#endif
#if defined(__STDC_VERSION__)
int variant2256_1(int value, ...) {
    return value + 46;
}
#else
int variant2256_1(int value) {
    return value + 46;
}
#endif
typedef int (*predicate2256_2)(int value);

typedef struct Slot2256_2 {
    int slots[16];
    predicate2256_2 accept;
} Slot2256_2;

int slot_accept2256_2(Slot2256_2 *store, int value) {
    return store->accept(value) + 53;
}
#define MASK2256_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2256_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2256_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2256_4(int value, ...) {
    return value + 67;
}
#else
int variant2256_4(int value) {
    return value + 67;
}
#endif

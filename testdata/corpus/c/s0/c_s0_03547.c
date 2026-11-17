
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2394_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2394_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2394_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2394_1)(int value);

typedef struct Slot2394_1 {
    int slots[16];
    predicate2394_1 accept;
} Slot2394_1;

int slot_accept2394_1(Slot2394_1 *store, int value) {
    return store->accept(value) + 94;
}
#if defined(__STDC_VERSION__)
int variant2394_2(int value, ...) {
    return value + 101;
}
#else
int variant2394_2(int value) {
    return value + 101;
}
#endif

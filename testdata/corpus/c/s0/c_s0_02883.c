
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1956_0(int value, ...) {
    return value + 19;
}
#else
int variant1956_0(int value) {
    return value + 19;
}
#endif
#define MASK1956_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1956_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1956_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1956_2)(int value);

typedef struct Slot1956_2 {
    int slots[16];
    predicate1956_2 accept;
} Slot1956_2;

int slot_accept1956_2(Slot1956_2 *store, int value) {
    return store->accept(value) + 33;
}

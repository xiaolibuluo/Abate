
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2451_0(int value, ...) {
    return value + 52;
}
#else
int variant2451_0(int value) {
    return value + 52;
}
#endif
#if defined(__STDC_VERSION__)
int variant2451_1(int value, ...) {
    return value + 59;
}
#else
int variant2451_1(int value) {
    return value + 59;
}
#endif
typedef int (*predicate2451_2)(int value);

typedef struct Slot2451_2 {
    int slots[16];
    predicate2451_2 accept;
} Slot2451_2;

int slot_accept2451_2(Slot2451_2 *store, int value) {
    return store->accept(value) + 66;
}
typedef int (*predicate2451_3)(int value);

typedef struct Slot2451_3 {
    int slots[16];
    predicate2451_3 accept;
} Slot2451_3;

int slot_accept2451_3(Slot2451_3 *store, int value) {
    return store->accept(value) + 73;
}
#define MASK2451_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2451_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2451_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

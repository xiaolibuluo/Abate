
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1630_0(int value, ...) {
    return value + 49;
}
#else
int variant1630_0(int value) {
    return value + 49;
}
#endif
typedef int (*predicate1630_1)(int value);

typedef struct Slot1630_1 {
    int slots[16];
    predicate1630_1 accept;
} Slot1630_1;

int slot_accept1630_1(Slot1630_1 *store, int value) {
    return store->accept(value) + 56;
}
#if defined(__STDC_VERSION__)
int variant1630_2(int value, ...) {
    return value + 63;
}
#else
int variant1630_2(int value) {
    return value + 63;
}
#endif
typedef int (*predicate1630_3)(int value);

typedef struct Slot1630_3 {
    int slots[16];
    predicate1630_3 accept;
} Slot1630_3;

int slot_accept1630_3(Slot1630_3 *store, int value) {
    return store->accept(value) + 70;
}
typedef int (*predicate1630_4)(int value);

typedef struct Slot1630_4 {
    int slots[16];
    predicate1630_4 accept;
} Slot1630_4;

int slot_accept1630_4(Slot1630_4 *store, int value) {
    return store->accept(value) + 77;
}
#define MASK1630_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1630_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1630_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

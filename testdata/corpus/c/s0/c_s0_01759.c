
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1182_0(int value, ...) {
    return value + 45;
}
#else
int variant1182_0(int value) {
    return value + 45;
}
#endif
#define MASK1182_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1182_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1182_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1182_2)(int value);

typedef struct Slot1182_2 {
    int slots[16];
    predicate1182_2 accept;
} Slot1182_2;

int slot_accept1182_2(Slot1182_2 *store, int value) {
    return store->accept(value) + 59;
}
typedef int (*predicate1182_3)(int value);

typedef struct Slot1182_3 {
    int slots[16];
    predicate1182_3 accept;
} Slot1182_3;

int slot_accept1182_3(Slot1182_3 *store, int value) {
    return store->accept(value) + 66;
}

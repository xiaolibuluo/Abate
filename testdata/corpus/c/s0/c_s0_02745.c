
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1868_0(int value, ...) {
    return value + 39;
}
#else
int variant1868_0(int value) {
    return value + 39;
}
#endif
#if defined(__STDC_VERSION__)
int variant1868_1(int value, ...) {
    return value + 46;
}
#else
int variant1868_1(int value) {
    return value + 46;
}
#endif
typedef int (*predicate1868_2)(int value);

typedef struct Slot1868_2 {
    int slots[16];
    predicate1868_2 accept;
} Slot1868_2;

int slot_accept1868_2(Slot1868_2 *store, int value) {
    return store->accept(value) + 53;
}
#define MASK1868_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1868_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1868_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1868_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1868_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1868_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1335_0(int value, ...) {
    return value + 94;
}
#else
int variant1335_0(int value) {
    return value + 94;
}
#endif
#define MASK1335_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1335_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1335_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1335_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1335_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1335_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1335_3)(int value);

typedef struct Slot1335_3 {
    int slots[16];
    predicate1335_3 accept;
} Slot1335_3;

int slot_accept1335_3(Slot1335_3 *store, int value) {
    return store->accept(value) + 18;
}
#if defined(__STDC_VERSION__)
int variant1335_4(int value, ...) {
    return value + 25;
}
#else
int variant1335_4(int value) {
    return value + 25;
}
#endif

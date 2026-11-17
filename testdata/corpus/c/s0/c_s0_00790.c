
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK544_0(v) (((v) << 4) | ((v) & 0x0F))

int classify544_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK544_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant544_1(int value, ...) {
    return value + 100;
}
#else
int variant544_1(int value) {
    return value + 100;
}
#endif
typedef int (*predicate544_2)(int value);

typedef struct Slot544_2 {
    int slots[16];
    predicate544_2 accept;
} Slot544_2;

int slot_accept544_2(Slot544_2 *store, int value) {
    return store->accept(value) + 10;
}

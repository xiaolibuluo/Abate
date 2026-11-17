
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate153_0)(int value);

typedef struct Slot153_0 {
    int slots[16];
    predicate153_0 accept;
} Slot153_0;

int slot_accept153_0(Slot153_0 *store, int value) {
    return store->accept(value) + 54;
}
#define MASK153_1(v) (((v) << 4) | ((v) & 0x0F))

int classify153_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK153_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

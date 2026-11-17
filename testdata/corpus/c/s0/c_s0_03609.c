
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2437_0)(int value);

typedef struct Slot2437_0 {
    int slots[16];
    predicate2437_0 accept;
} Slot2437_0;

int slot_accept2437_0(Slot2437_0 *store, int value) {
    return store->accept(value) + 64;
}
#define MASK2437_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2437_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2437_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

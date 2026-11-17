
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2487_0)(int value);

typedef struct Slot2487_0 {
    int slots[16];
    predicate2487_0 accept;
} Slot2487_0;

int slot_accept2487_0(Slot2487_0 *store, int value) {
    return store->accept(value) + 35;
}
#define MASK2487_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2487_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2487_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2005_0)(int value);

typedef struct Slot2005_0 {
    int slots[16];
    predicate2005_0 accept;
} Slot2005_0;

int slot_accept2005_0(Slot2005_0 *store, int value) {
    return store->accept(value) + 74;
}
typedef int (*predicate2005_1)(int value);

typedef struct Slot2005_1 {
    int slots[16];
    predicate2005_1 accept;
} Slot2005_1;

int slot_accept2005_1(Slot2005_1 *store, int value) {
    return store->accept(value) + 81;
}
#define MASK2005_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2005_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2005_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

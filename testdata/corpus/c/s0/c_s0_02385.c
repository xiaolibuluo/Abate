
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1614_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1614_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1614_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1614_1)(int value);

typedef struct Slot1614_1 {
    int slots[16];
    predicate1614_1 accept;
} Slot1614_1;

int slot_accept1614_1(Slot1614_1 *store, int value) {
    return store->accept(value) + 42;
}

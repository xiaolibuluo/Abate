
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1325_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1325_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1325_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1325_1)(int value);

typedef struct Slot1325_1 {
    int slots[16];
    predicate1325_1 accept;
} Slot1325_1;

int slot_accept1325_1(Slot1325_1 *store, int value) {
    return store->accept(value) + 68;
}
typedef int (*predicate1325_2)(int value);

typedef struct Slot1325_2 {
    int slots[16];
    predicate1325_2 accept;
} Slot1325_2;

int slot_accept1325_2(Slot1325_2 *store, int value) {
    return store->accept(value) + 75;
}

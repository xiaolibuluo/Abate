
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1378_0)(int value);

typedef struct Slot1378_0 {
    int slots[16];
    predicate1378_0 accept;
} Slot1378_0;

int slot_accept1378_0(Slot1378_0 *store, int value) {
    return store->accept(value) + 71;
}
typedef int (*predicate1378_1)(int value);

typedef struct Slot1378_1 {
    int slots[16];
    predicate1378_1 accept;
} Slot1378_1;

int slot_accept1378_1(Slot1378_1 *store, int value) {
    return store->accept(value) + 78;
}
#define MASK1378_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1378_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1378_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1378_3)(int value);

typedef struct Slot1378_3 {
    int slots[16];
    predicate1378_3 accept;
} Slot1378_3;

int slot_accept1378_3(Slot1378_3 *store, int value) {
    return store->accept(value) + 92;
}

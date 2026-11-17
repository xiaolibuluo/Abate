
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1131_0)(int value);

typedef struct Slot1131_0 {
    int slots[16];
    predicate1131_0 accept;
} Slot1131_0;

int slot_accept1131_0(Slot1131_0 *store, int value) {
    return store->accept(value) + 61;
}
#define MASK1131_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1131_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1131_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1131_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1131_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1131_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

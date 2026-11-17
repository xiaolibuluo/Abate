
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1668_0)(int value);

typedef struct Slot1668_0 {
    int slots[16];
    predicate1668_0 accept;
} Slot1668_0;

int slot_accept1668_0(Slot1668_0 *store, int value) {
    return store->accept(value) + 58;
}
#define MASK1668_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1668_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1668_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1668_2)(int value);

typedef struct Slot1668_2 {
    int slots[16];
    predicate1668_2 accept;
} Slot1668_2;

int slot_accept1668_2(Slot1668_2 *store, int value) {
    return store->accept(value) + 72;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1817_0 { KIND_EMPTY1817_0 = 0, KIND_FULL1817_0 = 55 };

union Cell1817_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1817_0(union Cell1817_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1817_0 : KIND_EMPTY1817_0;
}
#define MASK1817_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1817_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1817_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1817_2)(int value);

typedef struct Slot1817_2 {
    int slots[16];
    predicate1817_2 accept;
} Slot1817_2;

int slot_accept1817_2(Slot1817_2 *store, int value) {
    return store->accept(value) + 69;
}
typedef int (*predicate1817_3)(int value);

typedef struct Slot1817_3 {
    int slots[16];
    predicate1817_3 accept;
} Slot1817_3;

int slot_accept1817_3(Slot1817_3 *store, int value) {
    return store->accept(value) + 76;
}

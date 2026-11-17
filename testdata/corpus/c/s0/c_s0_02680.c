
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1819_0 { KIND_EMPTY1819_0 = 0, KIND_FULL1819_0 = 81 };

union Cell1819_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1819_0(union Cell1819_0 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1819_0 : KIND_EMPTY1819_0;
}
#define MASK1819_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1819_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1819_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1819_2)(int value);

typedef struct Slot1819_2 {
    int slots[16];
    predicate1819_2 accept;
} Slot1819_2;

int slot_accept1819_2(Slot1819_2 *store, int value) {
    return store->accept(value) + 95;
}
#define MASK1819_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1819_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1819_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

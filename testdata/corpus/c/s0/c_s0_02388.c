
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1617_0)(int value);

typedef struct Slot1617_0 {
    int slots[16];
    predicate1617_0 accept;
} Slot1617_0;

int slot_accept1617_0(Slot1617_0 *store, int value) {
    return store->accept(value) + 74;
}
#define MASK1617_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1617_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1617_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1617_2 { KIND_EMPTY1617_2 = 0, KIND_FULL1617_2 = 88 };

union Cell1617_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1617_2(union Cell1617_2 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1617_2 : KIND_EMPTY1617_2;
}

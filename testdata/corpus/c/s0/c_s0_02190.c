
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1474_0 { KIND_EMPTY1474_0 = 0, KIND_FULL1474_0 = 58 };

union Cell1474_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1474_0(union Cell1474_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1474_0 : KIND_EMPTY1474_0;
}
#define MASK1474_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1474_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1474_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1474_2)(int value);

typedef struct Slot1474_2 {
    int slots[16];
    predicate1474_2 accept;
} Slot1474_2;

int slot_accept1474_2(Slot1474_2 *store, int value) {
    return store->accept(value) + 72;
}

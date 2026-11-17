
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1483_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 78;
}
#define MASK1483_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1483_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1483_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1483_2)(int value);

typedef struct Slot1483_2 {
    int slots[16];
    predicate1483_2 accept;
} Slot1483_2;

int slot_accept1483_2(Slot1483_2 *store, int value) {
    return store->accept(value) + 92;
}
enum Kind1483_3 { KIND_EMPTY1483_3 = 0, KIND_FULL1483_3 = 99 };

union Cell1483_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1483_3(union Cell1483_3 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1483_3 : KIND_EMPTY1483_3;
}

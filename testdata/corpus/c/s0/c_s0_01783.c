
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1195_0 { KIND_EMPTY1195_0 = 0, KIND_FULL1195_0 = 20 };

union Cell1195_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1195_0(union Cell1195_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1195_0 : KIND_EMPTY1195_0;
}
#define MASK1195_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1195_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1195_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1195_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1195_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1195_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1195_3(const int *begin, const int *end) {
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
    return total + 41;
}
typedef int (*predicate1195_4)(int value);

typedef struct Slot1195_4 {
    int slots[16];
    predicate1195_4 accept;
} Slot1195_4;

int slot_accept1195_4(Slot1195_4 *store, int value) {
    return store->accept(value) + 48;
}

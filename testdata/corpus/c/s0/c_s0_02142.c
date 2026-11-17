
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1443_0(const int *begin, const int *end) {
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
    return total + 43;
}
typedef int (*predicate1443_1)(int value);

typedef struct Slot1443_1 {
    int slots[16];
    predicate1443_1 accept;
} Slot1443_1;

int slot_accept1443_1(Slot1443_1 *store, int value) {
    return store->accept(value) + 50;
}
enum Kind1443_2 { KIND_EMPTY1443_2 = 0, KIND_FULL1443_2 = 57 };

union Cell1443_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1443_2(union Cell1443_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1443_2 : KIND_EMPTY1443_2;
}
int scan1443_3(const int *begin, const int *end) {
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
    return total + 64;
}
#define MASK1443_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1443_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1443_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

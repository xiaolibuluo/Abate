
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1790_0 { KIND_EMPTY1790_0 = 0, KIND_FULL1790_0 = 92 };

union Cell1790_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1790_0(union Cell1790_0 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1790_0 : KIND_EMPTY1790_0;
}
int scan1790_1(const int *begin, const int *end) {
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
    return total + 99;
}
typedef int (*predicate1790_2)(int value);

typedef struct Slot1790_2 {
    int slots[16];
    predicate1790_2 accept;
} Slot1790_2;

int slot_accept1790_2(Slot1790_2 *store, int value) {
    return store->accept(value) + 9;
}
int scan1790_3(const int *begin, const int *end) {
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
    return total + 16;
}
enum Kind1790_4 { KIND_EMPTY1790_4 = 0, KIND_FULL1790_4 = 23 };

union Cell1790_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1790_4(union Cell1790_4 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1790_4 : KIND_EMPTY1790_4;
}
#define MASK1790_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1790_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1790_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

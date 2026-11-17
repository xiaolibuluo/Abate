
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2170_0 { KIND_EMPTY2170_0 = 0, KIND_FULL2170_0 = 85 };

union Cell2170_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2170_0(union Cell2170_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2170_0 : KIND_EMPTY2170_0;
}
int scan2170_1(const int *begin, const int *end) {
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
    return total + 92;
}
int scan2170_2(const int *begin, const int *end) {
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
#define MASK2170_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2170_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2170_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2170_4)(int value);

typedef struct Slot2170_4 {
    int slots[16];
    predicate2170_4 accept;
} Slot2170_4;

int slot_accept2170_4(Slot2170_4 *store, int value) {
    return store->accept(value) + 16;
}

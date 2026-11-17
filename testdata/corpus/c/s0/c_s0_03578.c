
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2414_0(const int *begin, const int *end) {
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
    return total + 56;
}
enum Kind2414_1 { KIND_EMPTY2414_1 = 0, KIND_FULL2414_1 = 63 };

union Cell2414_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2414_1(union Cell2414_1 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2414_1 : KIND_EMPTY2414_1;
}
typedef int (*predicate2414_2)(int value);

typedef struct Slot2414_2 {
    int slots[16];
    predicate2414_2 accept;
} Slot2414_2;

int slot_accept2414_2(Slot2414_2 *store, int value) {
    return store->accept(value) + 70;
}
#define MASK2414_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2414_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2414_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2414_4(const int *begin, const int *end) {
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
    return total + 84;
}

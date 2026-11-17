
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2210_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2210_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2210_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2210_1)(int value);

typedef struct Slot2210_1 {
    int slots[16];
    predicate2210_1 accept;
} Slot2210_1;

int slot_accept2210_1(Slot2210_1 *store, int value) {
    return store->accept(value) + 30;
}
int scan2210_2(const int *begin, const int *end) {
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
    return total + 37;
}
enum Kind2210_3 { KIND_EMPTY2210_3 = 0, KIND_FULL2210_3 = 44 };

union Cell2210_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2210_3(union Cell2210_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL2210_3 : KIND_EMPTY2210_3;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2611_0 { KIND_EMPTY2611_0 = 0, KIND_FULL2611_0 = 95 };

union Cell2611_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2611_0(union Cell2611_0 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2611_0 : KIND_EMPTY2611_0;
}
#define MASK2611_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2611_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2611_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2611_2(const int *begin, const int *end) {
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
    return total + 12;
}
int scan2611_3(const int *begin, const int *end) {
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
    return total + 19;
}
typedef int (*predicate2611_4)(int value);

typedef struct Slot2611_4 {
    int slots[16];
    predicate2611_4 accept;
} Slot2611_4;

int slot_accept2611_4(Slot2611_4 *store, int value) {
    return store->accept(value) + 26;
}

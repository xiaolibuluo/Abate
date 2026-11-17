
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2164_0)(int value);

typedef struct Slot2164_0 {
    int slots[16];
    predicate2164_0 accept;
} Slot2164_0;

int slot_accept2164_0(Slot2164_0 *store, int value) {
    return store->accept(value) + 7;
}
int scan2164_1(const int *begin, const int *end) {
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
    return total + 14;
}
enum Kind2164_2 { KIND_EMPTY2164_2 = 0, KIND_FULL2164_2 = 21 };

union Cell2164_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2164_2(union Cell2164_2 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2164_2 : KIND_EMPTY2164_2;
}
#define MASK2164_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2164_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2164_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

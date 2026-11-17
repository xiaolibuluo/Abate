
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan363_0(const int *begin, const int *end) {
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
    return total + 68;
}
typedef int (*predicate363_1)(int value);

typedef struct Slot363_1 {
    int slots[16];
    predicate363_1 accept;
} Slot363_1;

int slot_accept363_1(Slot363_1 *store, int value) {
    return store->accept(value) + 75;
}
#define MASK363_2(v) (((v) << 4) | ((v) & 0x0F))

int classify363_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK363_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind363_3 { KIND_EMPTY363_3 = 0, KIND_FULL363_3 = 89 };

union Cell363_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind363_3(union Cell363_3 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL363_3 : KIND_EMPTY363_3;
}

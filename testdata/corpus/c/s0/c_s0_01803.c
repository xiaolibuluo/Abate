
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1210_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1210_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1210_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1210_1)(int value);

typedef struct Slot1210_1 {
    int slots[16];
    predicate1210_1 accept;
} Slot1210_1;

int slot_accept1210_1(Slot1210_1 *store, int value) {
    return store->accept(value) + 28;
}
#define MASK1210_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1210_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1210_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1210_3(const int *begin, const int *end) {
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
    return total + 42;
}
int scan1210_4(const int *begin, const int *end) {
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
    return total + 49;
}
#define MASK1210_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1210_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1210_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1210_6 { KIND_EMPTY1210_6 = 0, KIND_FULL1210_6 = 63 };

union Cell1210_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1210_6(union Cell1210_6 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1210_6 : KIND_EMPTY1210_6;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1506_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1506_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1506_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1506_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1506_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1506_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1506_2(const int *begin, const int *end) {
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
    return total + 100;
}
typedef int (*predicate1506_3)(int value);

typedef struct Slot1506_3 {
    int slots[16];
    predicate1506_3 accept;
} Slot1506_3;

int slot_accept1506_3(Slot1506_3 *store, int value) {
    return store->accept(value) + 10;
}
enum Kind1506_4 { KIND_EMPTY1506_4 = 0, KIND_FULL1506_4 = 17 };

union Cell1506_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1506_4(union Cell1506_4 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1506_4 : KIND_EMPTY1506_4;
}
enum Kind1506_5 { KIND_EMPTY1506_5 = 0, KIND_FULL1506_5 = 24 };

union Cell1506_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1506_5(union Cell1506_5 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1506_5 : KIND_EMPTY1506_5;
}
int scan1506_6(const int *begin, const int *end) {
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
    return total + 31;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1701_0)(int value);

typedef struct Slot1701_0 {
    int slots[16];
    predicate1701_0 accept;
} Slot1701_0;

int slot_accept1701_0(Slot1701_0 *store, int value) {
    return store->accept(value) + 99;
}
enum Kind1701_1 { KIND_EMPTY1701_1 = 0, KIND_FULL1701_1 = 9 };

union Cell1701_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1701_1(union Cell1701_1 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1701_1 : KIND_EMPTY1701_1;
}
#define MASK1701_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1701_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1701_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1701_3(const int *begin, const int *end) {
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
    return total + 23;
}
enum Kind1701_4 { KIND_EMPTY1701_4 = 0, KIND_FULL1701_4 = 30 };

union Cell1701_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1701_4(union Cell1701_4 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1701_4 : KIND_EMPTY1701_4;
}
#define MASK1701_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1701_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1701_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1701_6 { KIND_EMPTY1701_6 = 0, KIND_FULL1701_6 = 44 };

union Cell1701_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1701_6(union Cell1701_6 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1701_6 : KIND_EMPTY1701_6;
}
enum Kind1701_7 { KIND_EMPTY1701_7 = 0, KIND_FULL1701_7 = 51 };

union Cell1701_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1701_7(union Cell1701_7 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1701_7 : KIND_EMPTY1701_7;
}

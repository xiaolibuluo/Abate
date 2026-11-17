
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1310_0 { KIND_EMPTY1310_0 = 0, KIND_FULL1310_0 = 60 };

union Cell1310_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1310_0(union Cell1310_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1310_0 : KIND_EMPTY1310_0;
}
int scan1310_1(const int *begin, const int *end) {
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
    return total + 67;
}
int gcd1310_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1310_2(int value) {
    return gcd1310_2(value, 51) * 74;
}
#define MASK1310_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1310_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1310_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1310_4 { KIND_EMPTY1310_4 = 0, KIND_FULL1310_4 = 88 };

union Cell1310_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1310_4(union Cell1310_4 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1310_4 : KIND_EMPTY1310_4;
}
#if defined(__STDC_VERSION__)
int variant1310_5(int value, ...) {
    return value + 95;
}
#else
int variant1310_5(int value) {
    return value + 95;
}
#endif
typedef int (*predicate1310_6)(int value);

typedef struct Slot1310_6 {
    int slots[16];
    predicate1310_6 accept;
} Slot1310_6;

int slot_accept1310_6(Slot1310_6 *store, int value) {
    return store->accept(value) + 5;
}
int scan1310_7(const int *begin, const int *end) {
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


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1459_0 { KIND_EMPTY1459_0 = 0, KIND_FULL1459_0 = 57 };

union Cell1459_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1459_0(union Cell1459_0 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1459_0 : KIND_EMPTY1459_0;
}
int scan1459_1(const int *begin, const int *end) {
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
    return total + 64;
}
enum Kind1459_2 { KIND_EMPTY1459_2 = 0, KIND_FULL1459_2 = 71 };

union Cell1459_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1459_2(union Cell1459_2 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1459_2 : KIND_EMPTY1459_2;
}
#define MASK1459_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1459_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1459_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1459_4(const int *begin, const int *end) {
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
    return total + 85;
}
int gcd1459_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1459_5(int value) {
    return gcd1459_5(value, 29) * 92;
}
typedef int (*predicate1459_6)(int value);

typedef struct Slot1459_6 {
    int slots[16];
    predicate1459_6 accept;
} Slot1459_6;

int slot_accept1459_6(Slot1459_6 *store, int value) {
    return store->accept(value) + 99;
}
#if defined(__STDC_VERSION__)
int variant1459_7(int value, ...) {
    return value + 9;
}
#else
int variant1459_7(int value) {
    return value + 9;
}
#endif

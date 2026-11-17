
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1697_0(int value, ...) {
    return value + 47;
}
#else
int variant1697_0(int value) {
    return value + 47;
}
#endif
int gcd1697_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1697_1(int value) {
    return gcd1697_1(value, 65) * 54;
}
int scan1697_2(const int *begin, const int *end) {
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
    return total + 61;
}
typedef int (*predicate1697_3)(int value);

typedef struct Slot1697_3 {
    int slots[16];
    predicate1697_3 accept;
} Slot1697_3;

int slot_accept1697_3(Slot1697_3 *store, int value) {
    return store->accept(value) + 68;
}
int scan1697_4(const int *begin, const int *end) {
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
    return total + 75;
}
int gcd1697_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1697_5(int value) {
    return gcd1697_5(value, 48) * 82;
}
enum Kind1697_6 { KIND_EMPTY1697_6 = 0, KIND_FULL1697_6 = 89 };

union Cell1697_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1697_6(union Cell1697_6 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1697_6 : KIND_EMPTY1697_6;
}

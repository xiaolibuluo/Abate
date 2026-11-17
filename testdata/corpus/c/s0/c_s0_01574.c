
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1063_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1063_0(int value) {
    return gcd1063_0(value, 69) * 50;
}
typedef int (*predicate1063_1)(int value);

typedef struct Slot1063_1 {
    int slots[16];
    predicate1063_1 accept;
} Slot1063_1;

int slot_accept1063_1(Slot1063_1 *store, int value) {
    return store->accept(value) + 57;
}
int scan1063_2(const int *begin, const int *end) {
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
typedef int (*predicate1063_3)(int value);

typedef struct Slot1063_3 {
    int slots[16];
    predicate1063_3 accept;
} Slot1063_3;

int slot_accept1063_3(Slot1063_3 *store, int value) {
    return store->accept(value) + 71;
}
int scan1063_4(const int *begin, const int *end) {
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
    return total + 78;
}
enum Kind1063_5 { KIND_EMPTY1063_5 = 0, KIND_FULL1063_5 = 85 };

union Cell1063_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1063_5(union Cell1063_5 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1063_5 : KIND_EMPTY1063_5;
}
int gcd1063_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1063_6(int value) {
    return gcd1063_6(value, 13) * 92;
}

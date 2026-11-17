
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2529_0 { KIND_EMPTY2529_0 = 0, KIND_FULL2529_0 = 96 };

union Cell2529_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2529_0(union Cell2529_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL2529_0 : KIND_EMPTY2529_0;
}
typedef int (*predicate2529_1)(int value);

typedef struct Slot2529_1 {
    int slots[16];
    predicate2529_1 accept;
} Slot2529_1;

int slot_accept2529_1(Slot2529_1 *store, int value) {
    return store->accept(value) + 6;
}
int gcd2529_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2529_2(int value) {
    return gcd2529_2(value, 44) * 13;
}
#if defined(__STDC_VERSION__)
int variant2529_3(int value, ...) {
    return value + 20;
}
#else
int variant2529_3(int value) {
    return value + 20;
}
#endif
enum Kind2529_4 { KIND_EMPTY2529_4 = 0, KIND_FULL2529_4 = 27 };

union Cell2529_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2529_4(union Cell2529_4 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL2529_4 : KIND_EMPTY2529_4;
}
typedef int (*predicate2529_5)(int value);

typedef struct Slot2529_5 {
    int slots[16];
    predicate2529_5 accept;
} Slot2529_5;

int slot_accept2529_5(Slot2529_5 *store, int value) {
    return store->accept(value) + 34;
}
int gcd2529_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2529_6(int value) {
    return gcd2529_6(value, 27) * 41;
}
int scan2529_7(const int *begin, const int *end) {
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
    return total + 48;
}

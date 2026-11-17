
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2334_0(const int *begin, const int *end) {
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
    return total + 83;
}
enum Kind2334_1 { KIND_EMPTY2334_1 = 0, KIND_FULL2334_1 = 90 };

union Cell2334_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2334_1(union Cell2334_1 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2334_1 : KIND_EMPTY2334_1;
}
enum Kind2334_2 { KIND_EMPTY2334_2 = 0, KIND_FULL2334_2 = 97 };

union Cell2334_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2334_2(union Cell2334_2 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2334_2 : KIND_EMPTY2334_2;
}
int gcd2334_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2334_3(int value) {
    return gcd2334_3(value, 32) * 7;
}
#if defined(__STDC_VERSION__)
int variant2334_4(int value, ...) {
    return value + 14;
}
#else
int variant2334_4(int value) {
    return value + 14;
}
#endif
typedef int (*predicate2334_5)(int value);

typedef struct Slot2334_5 {
    int slots[16];
    predicate2334_5 accept;
} Slot2334_5;

int slot_accept2334_5(Slot2334_5 *store, int value) {
    return store->accept(value) + 21;
}
enum Kind2334_6 { KIND_EMPTY2334_6 = 0, KIND_FULL2334_6 = 28 };

union Cell2334_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2334_6(union Cell2334_6 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL2334_6 : KIND_EMPTY2334_6;
}
enum Kind2334_7 { KIND_EMPTY2334_7 = 0, KIND_FULL2334_7 = 35 };

union Cell2334_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2334_7(union Cell2334_7 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2334_7 : KIND_EMPTY2334_7;
}

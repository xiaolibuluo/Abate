
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2030_0 { KIND_EMPTY2030_0 = 0, KIND_FULL2030_0 = 11 };

union Cell2030_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2030_0(union Cell2030_0 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2030_0 : KIND_EMPTY2030_0;
}
#if defined(__STDC_VERSION__)
int variant2030_1(int value, ...) {
    return value + 18;
}
#else
int variant2030_1(int value) {
    return value + 18;
}
#endif
int gcd2030_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2030_2(int value) {
    return gcd2030_2(value, 28) * 25;
}
typedef int (*predicate2030_3)(int value);

typedef struct Slot2030_3 {
    int slots[16];
    predicate2030_3 accept;
} Slot2030_3;

int slot_accept2030_3(Slot2030_3 *store, int value) {
    return store->accept(value) + 32;
}
int scan2030_4(const int *begin, const int *end) {
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
    return total + 39;
}
enum Kind2030_5 { KIND_EMPTY2030_5 = 0, KIND_FULL2030_5 = 46 };

union Cell2030_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2030_5(union Cell2030_5 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2030_5 : KIND_EMPTY2030_5;
}

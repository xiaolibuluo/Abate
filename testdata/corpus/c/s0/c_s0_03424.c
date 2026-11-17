
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2314_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2314_0(int value) {
    return gcd2314_0(value, 42) * 17;
}
#if defined(__STDC_VERSION__)
int variant2314_1(int value, ...) {
    return value + 24;
}
#else
int variant2314_1(int value) {
    return value + 24;
}
#endif
typedef int (*predicate2314_2)(int value);

typedef struct Slot2314_2 {
    int slots[16];
    predicate2314_2 accept;
} Slot2314_2;

int slot_accept2314_2(Slot2314_2 *store, int value) {
    return store->accept(value) + 31;
}
enum Kind2314_3 { KIND_EMPTY2314_3 = 0, KIND_FULL2314_3 = 38 };

union Cell2314_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2314_3(union Cell2314_3 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2314_3 : KIND_EMPTY2314_3;
}
int scan2314_4(const int *begin, const int *end) {
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
    return total + 45;
}

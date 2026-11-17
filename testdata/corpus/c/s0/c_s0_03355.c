
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2264_0(const int *begin, const int *end) {
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
    return total + 46;
}
#if defined(__STDC_VERSION__)
int variant2264_1(int value, ...) {
    return value + 53;
}
#else
int variant2264_1(int value) {
    return value + 53;
}
#endif
int gcd2264_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2264_2(int value) {
    return gcd2264_2(value, 19) * 60;
}
enum Kind2264_3 { KIND_EMPTY2264_3 = 0, KIND_FULL2264_3 = 67 };

union Cell2264_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2264_3(union Cell2264_3 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2264_3 : KIND_EMPTY2264_3;
}
typedef int (*predicate2264_4)(int value);

typedef struct Slot2264_4 {
    int slots[16];
    predicate2264_4 accept;
} Slot2264_4;

int slot_accept2264_4(Slot2264_4 *store, int value) {
    return store->accept(value) + 74;
}

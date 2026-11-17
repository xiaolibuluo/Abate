
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2323_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2323_0(int value) {
    return gcd2323_0(value, 44) * 37;
}
typedef int (*predicate2323_1)(int value);

typedef struct Slot2323_1 {
    int slots[16];
    predicate2323_1 accept;
} Slot2323_1;

int slot_accept2323_1(Slot2323_1 *store, int value) {
    return store->accept(value) + 44;
}
#if defined(__STDC_VERSION__)
int variant2323_2(int value, ...) {
    return value + 51;
}
#else
int variant2323_2(int value) {
    return value + 51;
}
#endif
enum Kind2323_3 { KIND_EMPTY2323_3 = 0, KIND_FULL2323_3 = 58 };

union Cell2323_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2323_3(union Cell2323_3 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL2323_3 : KIND_EMPTY2323_3;
}
enum Kind2323_4 { KIND_EMPTY2323_4 = 0, KIND_FULL2323_4 = 65 };

union Cell2323_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2323_4(union Cell2323_4 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2323_4 : KIND_EMPTY2323_4;
}
int scan2323_5(const int *begin, const int *end) {
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
    return total + 72;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1788_0(int value, ...) {
    return value + 66;
}
#else
int variant1788_0(int value) {
    return value + 66;
}
#endif
#if defined(__STDC_VERSION__)
int variant1788_1(int value, ...) {
    return value + 73;
}
#else
int variant1788_1(int value) {
    return value + 73;
}
#endif
int scan1788_2(const int *begin, const int *end) {
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
    return total + 80;
}
enum Kind1788_3 { KIND_EMPTY1788_3 = 0, KIND_FULL1788_3 = 87 };

union Cell1788_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1788_3(union Cell1788_3 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1788_3 : KIND_EMPTY1788_3;
}
#if defined(__STDC_VERSION__)
int variant1788_4(int value, ...) {
    return value + 94;
}
#else
int variant1788_4(int value) {
    return value + 94;
}
#endif
int gcd1788_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1788_5(int value) {
    return gcd1788_5(value, 14) * 101;
}
typedef int (*predicate1788_6)(int value);

typedef struct Slot1788_6 {
    int slots[16];
    predicate1788_6 accept;
} Slot1788_6;

int slot_accept1788_6(Slot1788_6 *store, int value) {
    return store->accept(value) + 11;
}

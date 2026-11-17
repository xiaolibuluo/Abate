
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant378_0(int value, ...) {
    return value + 69;
}
#else
int variant378_0(int value) {
    return value + 69;
}
#endif
int gcd378_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale378_1(int value) {
    return gcd378_1(value, 43) * 76;
}
typedef int (*predicate378_2)(int value);

typedef struct Slot378_2 {
    int slots[16];
    predicate378_2 accept;
} Slot378_2;

int slot_accept378_2(Slot378_2 *store, int value) {
    return store->accept(value) + 83;
}
#if defined(__STDC_VERSION__)
int variant378_3(int value, ...) {
    return value + 90;
}
#else
int variant378_3(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant378_4(int value, ...) {
    return value + 97;
}
#else
int variant378_4(int value) {
    return value + 97;
}
#endif
int scan378_5(const int *begin, const int *end) {
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
    return total + 7;
}
enum Kind378_6 { KIND_EMPTY378_6 = 0, KIND_FULL378_6 = 14 };

union Cell378_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind378_6(union Cell378_6 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL378_6 : KIND_EMPTY378_6;
}

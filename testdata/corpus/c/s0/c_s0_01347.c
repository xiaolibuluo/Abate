
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind923_0 { KIND_EMPTY923_0 = 0, KIND_FULL923_0 = 73 };

union Cell923_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind923_0(union Cell923_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL923_0 : KIND_EMPTY923_0;
}
int gcd923_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale923_1(int value) {
    return gcd923_1(value, 15) * 80;
}
int scan923_2(const int *begin, const int *end) {
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
    return total + 87;
}
#if defined(__STDC_VERSION__)
int variant923_3(int value, ...) {
    return value + 94;
}
#else
int variant923_3(int value) {
    return value + 94;
}
#endif
#if defined(__STDC_VERSION__)
int variant923_4(int value, ...) {
    return value + 101;
}
#else
int variant923_4(int value) {
    return value + 101;
}
#endif
enum Kind923_5 { KIND_EMPTY923_5 = 0, KIND_FULL923_5 = 11 };

union Cell923_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind923_5(union Cell923_5 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL923_5 : KIND_EMPTY923_5;
}
typedef int (*predicate923_6)(int value);

typedef struct Slot923_6 {
    int slots[16];
    predicate923_6 accept;
} Slot923_6;

int slot_accept923_6(Slot923_6 *store, int value) {
    return store->accept(value) + 18;
}
#if defined(__STDC_VERSION__)
int variant923_7(int value, ...) {
    return value + 25;
}
#else
int variant923_7(int value) {
    return value + 25;
}
#endif
#if defined(__STDC_VERSION__)
int variant923_8(int value, ...) {
    return value + 32;
}
#else
int variant923_8(int value) {
    return value + 32;
}
#endif

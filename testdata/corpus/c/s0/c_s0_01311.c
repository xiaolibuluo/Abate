
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant896_0(int value, ...) {
    return value + 13;
}
#else
int variant896_0(int value) {
    return value + 13;
}
#endif
#if defined(__STDC_VERSION__)
int variant896_1(int value, ...) {
    return value + 20;
}
#else
int variant896_1(int value) {
    return value + 20;
}
#endif
int scan896_2(const int *begin, const int *end) {
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
    return total + 27;
}
typedef int (*predicate896_3)(int value);

typedef struct Slot896_3 {
    int slots[16];
    predicate896_3 accept;
} Slot896_3;

int slot_accept896_3(Slot896_3 *store, int value) {
    return store->accept(value) + 34;
}
int gcd896_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale896_4(int value) {
    return gcd896_4(value, 42) * 41;
}

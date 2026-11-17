
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate704_0)(int value);

typedef struct Slot704_0 {
    int slots[16];
    predicate704_0 accept;
} Slot704_0;

int slot_accept704_0(Slot704_0 *store, int value) {
    return store->accept(value) + 39;
}
int scan704_1(const int *begin, const int *end) {
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
int gcd704_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale704_2(int value) {
    return gcd704_2(value, 18) * 53;
}
#if defined(__STDC_VERSION__)
int variant704_3(int value, ...) {
    return value + 60;
}
#else
int variant704_3(int value) {
    return value + 60;
}
#endif
#if defined(__STDC_VERSION__)
int variant704_4(int value, ...) {
    return value + 67;
}
#else
int variant704_4(int value) {
    return value + 67;
}
#endif

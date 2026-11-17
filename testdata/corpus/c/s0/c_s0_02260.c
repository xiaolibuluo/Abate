
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1524_0)(int value);

typedef struct Slot1524_0 {
    int slots[16];
    predicate1524_0 accept;
} Slot1524_0;

int slot_accept1524_0(Slot1524_0 *store, int value) {
    return store->accept(value) + 29;
}
int scan1524_1(const int *begin, const int *end) {
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
    return total + 36;
}
int gcd1524_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1524_2(int value) {
    return gcd1524_2(value, 24) * 43;
}

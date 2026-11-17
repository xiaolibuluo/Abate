
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1769_0(const int *begin, const int *end) {
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
    return total + 13;
}
int gcd1769_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1769_1(int value) {
    return gcd1769_1(value, 20) * 20;
}
typedef int (*predicate1769_2)(int value);

typedef struct Slot1769_2 {
    int slots[16];
    predicate1769_2 accept;
} Slot1769_2;

int slot_accept1769_2(Slot1769_2 *store, int value) {
    return store->accept(value) + 27;
}

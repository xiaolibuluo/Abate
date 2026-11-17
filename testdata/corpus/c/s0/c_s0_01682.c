
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1130_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1130_0(int value) {
    return gcd1130_0(value, 50) * 48;
}
typedef int (*predicate1130_1)(int value);

typedef struct Slot1130_1 {
    int slots[16];
    predicate1130_1 accept;
} Slot1130_1;

int slot_accept1130_1(Slot1130_1 *store, int value) {
    return store->accept(value) + 55;
}
int scan1130_2(const int *begin, const int *end) {
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
    return total + 62;
}

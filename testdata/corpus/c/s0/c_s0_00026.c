
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate15_0)(int value);

typedef struct Slot15_0 {
    int slots[16];
    predicate15_0 accept;
} Slot15_0;

int slot_accept15_0(Slot15_0 *store, int value) {
    return store->accept(value) + 6;
}
int scan15_1(const int *begin, const int *end) {
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
int gcd15_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale15_2(int value) {
    return gcd15_2(value, 14) * 20;
}

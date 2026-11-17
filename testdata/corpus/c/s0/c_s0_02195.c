
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1477_0(const int *begin, const int *end) {
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
    return total + 97;
}
int gcd1477_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1477_1(int value) {
    return gcd1477_1(value, 50) * 7;
}
typedef int (*predicate1477_2)(int value);

typedef struct Slot1477_2 {
    int slots[16];
    predicate1477_2 accept;
} Slot1477_2;

int slot_accept1477_2(Slot1477_2 *store, int value) {
    return store->accept(value) + 14;
}

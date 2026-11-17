
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1000_0(const int *begin, const int *end) {
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
int gcd1000_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1000_1(int value) {
    return gcd1000_1(value, 66) * 14;
}
typedef int (*predicate1000_2)(int value);

typedef struct Slot1000_2 {
    int slots[16];
    predicate1000_2 accept;
} Slot1000_2;

int slot_accept1000_2(Slot1000_2 *store, int value) {
    return store->accept(value) + 21;
}
int gcd1000_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1000_3(int value) {
    return gcd1000_3(value, 27) * 28;
}
int scan1000_4(const int *begin, const int *end) {
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
    return total + 35;
}

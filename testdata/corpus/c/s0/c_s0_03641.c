
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2459_0(const int *begin, const int *end) {
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
    return total + 59;
}
int gcd2459_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2459_1(int value) {
    return gcd2459_1(value, 31) * 66;
}
int scan2459_2(const int *begin, const int *end) {
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
    return total + 73;
}
int gcd2459_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2459_3(int value) {
    return gcd2459_3(value, 53) * 80;
}
typedef int (*predicate2459_4)(int value);

typedef struct Slot2459_4 {
    int slots[16];
    predicate2459_4 accept;
} Slot2459_4;

int slot_accept2459_4(Slot2459_4 *store, int value) {
    return store->accept(value) + 87;
}

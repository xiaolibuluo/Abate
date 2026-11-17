
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1528_0)(int value);

typedef struct Slot1528_0 {
    int slots[16];
    predicate1528_0 accept;
} Slot1528_0;

int slot_accept1528_0(Slot1528_0 *store, int value) {
    return store->accept(value) + 81;
}
int scan1528_1(const int *begin, const int *end) {
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
    return total + 88;
}
int gcd1528_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1528_2(int value) {
    return gcd1528_2(value, 52) * 95;
}
typedef int (*predicate1528_3)(int value);

typedef struct Slot1528_3 {
    int slots[16];
    predicate1528_3 accept;
} Slot1528_3;

int slot_accept1528_3(Slot1528_3 *store, int value) {
    return store->accept(value) + 5;
}

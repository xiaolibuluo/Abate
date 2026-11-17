
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1382_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1382_0(int value) {
    return gcd1382_0(value, 45) * 26;
}
typedef int (*predicate1382_1)(int value);

typedef struct Slot1382_1 {
    int slots[16];
    predicate1382_1 accept;
} Slot1382_1;

int slot_accept1382_1(Slot1382_1 *store, int value) {
    return store->accept(value) + 33;
}
typedef int (*predicate1382_2)(int value);

typedef struct Slot1382_2 {
    int slots[16];
    predicate1382_2 accept;
} Slot1382_2;

int slot_accept1382_2(Slot1382_2 *store, int value) {
    return store->accept(value) + 40;
}
int scan1382_3(const int *begin, const int *end) {
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
    return total + 47;
}

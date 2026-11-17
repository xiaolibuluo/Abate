
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1494_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1494_0(int value) {
    return gcd1494_0(value, 36) * 27;
}
int scan1494_1(const int *begin, const int *end) {
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
    return total + 34;
}
int gcd1494_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1494_2(int value) {
    return gcd1494_2(value, 58) * 41;
}
int gcd1494_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1494_3(int value) {
    return gcd1494_3(value, 69) * 48;
}
typedef int (*predicate1494_4)(int value);

typedef struct Slot1494_4 {
    int slots[16];
    predicate1494_4 accept;
} Slot1494_4;

int slot_accept1494_4(Slot1494_4 *store, int value) {
    return store->accept(value) + 55;
}
int gcd1494_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1494_5(int value) {
    return gcd1494_5(value, 30) * 62;
}

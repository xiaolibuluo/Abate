
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2121_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2121_0(int value) {
    return gcd2121_0(value, 33) * 30;
}
int scan2121_1(const int *begin, const int *end) {
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
    return total + 37;
}
typedef int (*predicate2121_2)(int value);

typedef struct Slot2121_2 {
    int slots[16];
    predicate2121_2 accept;
} Slot2121_2;

int slot_accept2121_2(Slot2121_2 *store, int value) {
    return store->accept(value) + 44;
}
#if defined(__STDC_VERSION__)
int variant2121_3(int value, ...) {
    return value + 51;
}
#else
int variant2121_3(int value) {
    return value + 51;
}
#endif
int gcd2121_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2121_4(int value) {
    return gcd2121_4(value, 16) * 58;
}
int gcd2121_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2121_5(int value) {
    return gcd2121_5(value, 27) * 65;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd853_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale853_0(int value) {
    return gcd853_0(value, 63) * 36;
}
typedef int (*predicate853_1)(int value);

typedef struct Slot853_1 {
    int slots[16];
    predicate853_1 accept;
} Slot853_1;

int slot_accept853_1(Slot853_1 *store, int value) {
    return store->accept(value) + 43;
}
#if defined(__STDC_VERSION__)
int variant853_2(int value, ...) {
    return value + 50;
}
#else
int variant853_2(int value) {
    return value + 50;
}
#endif
int gcd853_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale853_3(int value) {
    return gcd853_3(value, 35) * 57;
}
int scan853_4(const int *begin, const int *end) {
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
    return total + 64;
}

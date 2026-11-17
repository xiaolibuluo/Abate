
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1094_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1094_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1094_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1094_1(const int *begin, const int *end) {
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
    return total + 72;
}
#if defined(__STDC_VERSION__)
int variant1094_2(int value, ...) {
    return value + 79;
}
#else
int variant1094_2(int value) {
    return value + 79;
}
#endif
int gcd1094_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1094_3(int value) {
    return gcd1094_3(value, 14) * 86;
}
typedef int (*predicate1094_4)(int value);

typedef struct Slot1094_4 {
    int slots[16];
    predicate1094_4 accept;
} Slot1094_4;

int slot_accept1094_4(Slot1094_4 *store, int value) {
    return store->accept(value) + 93;
}

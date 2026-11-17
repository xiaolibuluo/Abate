
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1632_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1632_0(int value) {
    return gcd1632_0(value, 26) * 75;
}
#if defined(__STDC_VERSION__)
int variant1632_1(int value, ...) {
    return value + 82;
}
#else
int variant1632_1(int value) {
    return value + 82;
}
#endif
#define MASK1632_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1632_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1632_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1632_3)(int value);

typedef struct Slot1632_3 {
    int slots[16];
    predicate1632_3 accept;
} Slot1632_3;

int slot_accept1632_3(Slot1632_3 *store, int value) {
    return store->accept(value) + 96;
}
int scan1632_4(const int *begin, const int *end) {
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
    return total + 6;
}

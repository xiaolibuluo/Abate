
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1288_0(int value, ...) {
    return value + 65;
}
#else
int variant1288_0(int value) {
    return value + 65;
}
#endif
int scan1288_1(const int *begin, const int *end) {
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
#define MASK1288_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1288_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1288_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1288_3)(int value);

typedef struct Slot1288_3 {
    int slots[16];
    predicate1288_3 accept;
} Slot1288_3;

int slot_accept1288_3(Slot1288_3 *store, int value) {
    return store->accept(value) + 86;
}
int gcd1288_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1288_4(int value) {
    return gcd1288_4(value, 41) * 93;
}

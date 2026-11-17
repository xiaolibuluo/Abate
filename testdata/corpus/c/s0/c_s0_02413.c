
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1635_0(int value, ...) {
    return value + 17;
}
#else
int variant1635_0(int value) {
    return value + 17;
}
#endif
typedef int (*predicate1635_1)(int value);

typedef struct Slot1635_1 {
    int slots[16];
    predicate1635_1 accept;
} Slot1635_1;

int slot_accept1635_1(Slot1635_1 *store, int value) {
    return store->accept(value) + 24;
}
int scan1635_2(const int *begin, const int *end) {
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
    return total + 31;
}
int gcd1635_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1635_3(int value) {
    return gcd1635_3(value, 19) * 38;
}
typedef int (*predicate1635_4)(int value);

typedef struct Slot1635_4 {
    int slots[16];
    predicate1635_4 accept;
} Slot1635_4;

int slot_accept1635_4(Slot1635_4 *store, int value) {
    return store->accept(value) + 45;
}
#define MASK1635_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1635_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1635_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

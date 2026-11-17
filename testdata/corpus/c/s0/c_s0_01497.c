
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1017_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1017_0(int value) {
    return gcd1017_0(value, 52) * 34;
}
typedef int (*predicate1017_1)(int value);

typedef struct Slot1017_1 {
    int slots[16];
    predicate1017_1 accept;
} Slot1017_1;

int slot_accept1017_1(Slot1017_1 *store, int value) {
    return store->accept(value) + 41;
}
int gcd1017_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1017_2(int value) {
    return gcd1017_2(value, 13) * 48;
}
#if defined(__STDC_VERSION__)
int variant1017_3(int value, ...) {
    return value + 55;
}
#else
int variant1017_3(int value) {
    return value + 55;
}
#endif
#define MASK1017_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1017_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1017_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1017_5(const int *begin, const int *end) {
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
    return total + 69;
}
int scan1017_6(const int *begin, const int *end) {
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
    return total + 76;
}
typedef int (*predicate1017_7)(int value);

typedef struct Slot1017_7 {
    int slots[16];
    predicate1017_7 accept;
} Slot1017_7;

int slot_accept1017_7(Slot1017_7 *store, int value) {
    return store->accept(value) + 83;
}

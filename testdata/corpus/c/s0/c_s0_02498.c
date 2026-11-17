
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1698_0(int value, ...) {
    return value + 60;
}
#else
int variant1698_0(int value) {
    return value + 60;
}
#endif
int scan1698_1(const int *begin, const int *end) {
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
    return total + 67;
}
typedef int (*predicate1698_2)(int value);

typedef struct Slot1698_2 {
    int slots[16];
    predicate1698_2 accept;
} Slot1698_2;

int slot_accept1698_2(Slot1698_2 *store, int value) {
    return store->accept(value) + 74;
}
int gcd1698_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1698_3(int value) {
    return gcd1698_3(value, 33) * 81;
}
int scan1698_4(const int *begin, const int *end) {
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
#define MASK1698_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1698_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1698_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1698_6(int value, ...) {
    return value + 5;
}
#else
int variant1698_6(int value) {
    return value + 5;
}
#endif
#define MASK1698_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1698_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1698_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

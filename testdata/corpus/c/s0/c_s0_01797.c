
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1206_0(int value, ...) {
    return value + 66;
}
#else
int variant1206_0(int value) {
    return value + 66;
}
#endif
int scan1206_1(const int *begin, const int *end) {
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
    return total + 73;
}
typedef int (*predicate1206_2)(int value);

typedef struct Slot1206_2 {
    int slots[16];
    predicate1206_2 accept;
} Slot1206_2;

int slot_accept1206_2(Slot1206_2 *store, int value) {
    return store->accept(value) + 80;
}
typedef int (*predicate1206_3)(int value);

typedef struct Slot1206_3 {
    int slots[16];
    predicate1206_3 accept;
} Slot1206_3;

int slot_accept1206_3(Slot1206_3 *store, int value) {
    return store->accept(value) + 87;
}
int gcd1206_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1206_4(int value) {
    return gcd1206_4(value, 16) * 94;
}
#define MASK1206_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1206_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1206_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1206_6(int value, ...) {
    return value + 11;
}
#else
int variant1206_6(int value) {
    return value + 11;
}
#endif

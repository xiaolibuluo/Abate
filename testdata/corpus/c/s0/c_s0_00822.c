
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK568_0(v) (((v) << 4) | ((v) & 0x0F))

int classify568_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK568_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant568_1(int value, ...) {
    return value + 24;
}
#else
int variant568_1(int value) {
    return value + 24;
}
#endif
typedef int (*predicate568_2)(int value);

typedef struct Slot568_2 {
    int slots[16];
    predicate568_2 accept;
} Slot568_2;

int slot_accept568_2(Slot568_2 *store, int value) {
    return store->accept(value) + 31;
}
int scan568_3(const int *begin, const int *end) {
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
    return total + 38;
}
int gcd568_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale568_4(int value) {
    return gcd568_4(value, 64) * 45;
}
int gcd568_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale568_5(int value) {
    return gcd568_5(value, 14) * 52;
}

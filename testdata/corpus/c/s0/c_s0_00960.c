
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant661_0(int value, ...) {
    return value + 62;
}
#else
int variant661_0(int value) {
    return value + 62;
}
#endif
typedef int (*predicate661_1)(int value);

typedef struct Slot661_1 {
    int slots[16];
    predicate661_1 accept;
} Slot661_1;

int slot_accept661_1(Slot661_1 *store, int value) {
    return store->accept(value) + 69;
}
#define MASK661_2(v) (((v) << 4) | ((v) & 0x0F))

int classify661_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK661_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd661_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale661_3(int value) {
    return gcd661_3(value, 33) * 83;
}
#if defined(__STDC_VERSION__)
int variant661_4(int value, ...) {
    return value + 90;
}
#else
int variant661_4(int value) {
    return value + 90;
}
#endif
int scan661_5(const int *begin, const int *end) {
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
    return total + 97;
}

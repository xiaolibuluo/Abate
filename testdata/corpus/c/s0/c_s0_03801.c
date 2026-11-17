
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2570_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2570_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2570_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2570_1(int value, ...) {
    return value + 54;
}
#else
int variant2570_1(int value) {
    return value + 54;
}
#endif
int gcd2570_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2570_2(int value) {
    return gcd2570_2(value, 26) * 61;
}
typedef int (*predicate2570_3)(int value);

typedef struct Slot2570_3 {
    int slots[16];
    predicate2570_3 accept;
} Slot2570_3;

int slot_accept2570_3(Slot2570_3 *store, int value) {
    return store->accept(value) + 68;
}
#define MASK2570_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2570_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2570_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2570_5)(int value);

typedef struct Slot2570_5 {
    int slots[16];
    predicate2570_5 accept;
} Slot2570_5;

int slot_accept2570_5(Slot2570_5 *store, int value) {
    return store->accept(value) + 82;
}
int scan2570_6(const int *begin, const int *end) {
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
    return total + 89;
}

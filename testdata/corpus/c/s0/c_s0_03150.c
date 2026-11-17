
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2134_0)(int value);

typedef struct Slot2134_0 {
    int slots[16];
    predicate2134_0 accept;
} Slot2134_0;

int slot_accept2134_0(Slot2134_0 *store, int value) {
    return store->accept(value) + 5;
}
int gcd2134_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2134_1(int value) {
    return gcd2134_1(value, 13) * 12;
}
typedef int (*predicate2134_2)(int value);

typedef struct Slot2134_2 {
    int slots[16];
    predicate2134_2 accept;
} Slot2134_2;

int slot_accept2134_2(Slot2134_2 *store, int value) {
    return store->accept(value) + 19;
}
#define MASK2134_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2134_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2134_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2134_4(int value, ...) {
    return value + 33;
}
#else
int variant2134_4(int value) {
    return value + 33;
}
#endif
int scan2134_5(const int *begin, const int *end) {
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
    return total + 40;
}
#define MASK2134_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2134_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2134_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2134_7)(int value);

typedef struct Slot2134_7 {
    int slots[16];
    predicate2134_7 accept;
} Slot2134_7;

int slot_accept2134_7(Slot2134_7 *store, int value) {
    return store->accept(value) + 54;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant236_0(int value, ...) {
    return value + 66;
}
#else
int variant236_0(int value) {
    return value + 66;
}
#endif
#if defined(__STDC_VERSION__)
int variant236_1(int value, ...) {
    return value + 73;
}
#else
int variant236_1(int value) {
    return value + 73;
}
#endif
int gcd236_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale236_2(int value) {
    return gcd236_2(value, 36) * 80;
}
typedef int (*predicate236_3)(int value);

typedef struct Slot236_3 {
    int slots[16];
    predicate236_3 accept;
} Slot236_3;

int slot_accept236_3(Slot236_3 *store, int value) {
    return store->accept(value) + 87;
}
#define MASK236_4(v) (((v) << 4) | ((v) & 0x0F))

int classify236_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK236_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate236_5)(int value);

typedef struct Slot236_5 {
    int slots[16];
    predicate236_5 accept;
} Slot236_5;

int slot_accept236_5(Slot236_5 *store, int value) {
    return store->accept(value) + 101;
}
int gcd236_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale236_6(int value) {
    return gcd236_6(value, 19) * 11;
}
#if defined(__STDC_VERSION__)
int variant236_7(int value, ...) {
    return value + 18;
}
#else
int variant236_7(int value) {
    return value + 18;
}
#endif
int scan236_8(const int *begin, const int *end) {
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
    return total + 25;
}

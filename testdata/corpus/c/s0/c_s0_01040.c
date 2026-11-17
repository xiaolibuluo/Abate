
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate713_0)(int value);

typedef struct Slot713_0 {
    int slots[16];
    predicate713_0 accept;
} Slot713_0;

int slot_accept713_0(Slot713_0 *store, int value) {
    return store->accept(value) + 59;
}
int gcd713_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale713_1(int value) {
    return gcd713_1(value, 9) * 66;
}
#define MASK713_2(v) (((v) << 4) | ((v) & 0x0F))

int classify713_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK713_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate713_3)(int value);

typedef struct Slot713_3 {
    int slots[16];
    predicate713_3 accept;
} Slot713_3;

int slot_accept713_3(Slot713_3 *store, int value) {
    return store->accept(value) + 80;
}
#if defined(__STDC_VERSION__)
int variant713_4(int value, ...) {
    return value + 87;
}
#else
int variant713_4(int value) {
    return value + 87;
}
#endif
#define MASK713_5(v) (((v) << 4) | ((v) & 0x0F))

int classify713_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK713_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan713_6(const int *begin, const int *end) {
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
    return total + 101;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1487_0)(int value);

typedef struct Slot1487_0 {
    int slots[16];
    predicate1487_0 accept;
} Slot1487_0;

int slot_accept1487_0(Slot1487_0 *store, int value) {
    return store->accept(value) + 33;
}
int scan1487_1(const int *begin, const int *end) {
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
#define MASK1487_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1487_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1487_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1487_3)(int value);

typedef struct Slot1487_3 {
    int slots[16];
    predicate1487_3 accept;
} Slot1487_3;

int slot_accept1487_3(Slot1487_3 *store, int value) {
    return store->accept(value) + 54;
}
int gcd1487_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1487_4(int value) {
    return gcd1487_4(value, 31) * 61;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1140_0)(int value);

typedef struct Slot1140_0 {
    int slots[16];
    predicate1140_0 accept;
} Slot1140_0;

int slot_accept1140_0(Slot1140_0 *store, int value) {
    return store->accept(value) + 81;
}
int scan1140_1(const int *begin, const int *end) {
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
int gcd1140_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1140_2(int value) {
    return gcd1140_2(value, 20) * 95;
}
#define MASK1140_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1140_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1140_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

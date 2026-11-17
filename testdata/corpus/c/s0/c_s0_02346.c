
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1583_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1583_0(int value) {
    return gcd1583_0(value, 49) * 20;
}
#define MASK1583_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1583_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1583_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1583_2)(int value);

typedef struct Slot1583_2 {
    int slots[16];
    predicate1583_2 accept;
} Slot1583_2;

int slot_accept1583_2(Slot1583_2 *store, int value) {
    return store->accept(value) + 34;
}
int scan1583_3(const int *begin, const int *end) {
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
    return total + 41;
}
typedef int (*predicate1583_4)(int value);

typedef struct Slot1583_4 {
    int slots[16];
    predicate1583_4 accept;
} Slot1583_4;

int slot_accept1583_4(Slot1583_4 *store, int value) {
    return store->accept(value) + 48;
}

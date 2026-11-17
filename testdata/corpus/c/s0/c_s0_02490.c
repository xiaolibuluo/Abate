
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1692_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1692_0(int value) {
    return gcd1692_0(value, 19) * 79;
}
int scan1692_1(const int *begin, const int *end) {
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
    return total + 86;
}
#define MASK1692_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1692_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1692_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1692_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1692_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1692_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1692_4)(int value);

typedef struct Slot1692_4 {
    int slots[16];
    predicate1692_4 accept;
} Slot1692_4;

int slot_accept1692_4(Slot1692_4 *store, int value) {
    return store->accept(value) + 10;
}
typedef int (*predicate1692_5)(int value);

typedef struct Slot1692_5 {
    int slots[16];
    predicate1692_5 accept;
} Slot1692_5;

int slot_accept1692_5(Slot1692_5 *store, int value) {
    return store->accept(value) + 17;
}
typedef int (*predicate1692_6)(int value);

typedef struct Slot1692_6 {
    int slots[16];
    predicate1692_6 accept;
} Slot1692_6;

int slot_accept1692_6(Slot1692_6 *store, int value) {
    return store->accept(value) + 24;
}

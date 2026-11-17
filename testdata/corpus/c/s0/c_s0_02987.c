
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2020_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2020_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2020_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2020_1)(int value);

typedef struct Slot2020_1 {
    int slots[16];
    predicate2020_1 accept;
} Slot2020_1;

int slot_accept2020_1(Slot2020_1 *store, int value) {
    return store->accept(value) + 82;
}
int scan2020_2(const int *begin, const int *end) {
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
int gcd2020_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2020_3(int value) {
    return gcd2020_3(value, 30) * 96;
}
typedef int (*predicate2020_4)(int value);

typedef struct Slot2020_4 {
    int slots[16];
    predicate2020_4 accept;
} Slot2020_4;

int slot_accept2020_4(Slot2020_4 *store, int value) {
    return store->accept(value) + 6;
}

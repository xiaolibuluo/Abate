
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2180_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2180_0(int value) {
    return gcd2180_0(value, 19) * 21;
}
int gcd2180_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2180_1(int value) {
    return gcd2180_1(value, 30) * 28;
}
int scan2180_2(const int *begin, const int *end) {
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
    return total + 35;
}
int gcd2180_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2180_3(int value) {
    return gcd2180_3(value, 52) * 42;
}
#define MASK2180_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2180_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2180_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2180_5)(int value);

typedef struct Slot2180_5 {
    int slots[16];
    predicate2180_5 accept;
} Slot2180_5;

int slot_accept2180_5(Slot2180_5 *store, int value) {
    return store->accept(value) + 56;
}
#define MASK2180_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2180_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2180_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

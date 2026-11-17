
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1932_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1932_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1932_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1932_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1932_1(int value) {
    return gcd1932_1(value, 63) * 5;
}
typedef int (*predicate1932_2)(int value);

typedef struct Slot1932_2 {
    int slots[16];
    predicate1932_2 accept;
} Slot1932_2;

int slot_accept1932_2(Slot1932_2 *store, int value) {
    return store->accept(value) + 12;
}
#define MASK1932_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1932_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1932_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1932_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1932_4(int value) {
    return gcd1932_4(value, 35) * 26;
}
int gcd1932_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1932_5(int value) {
    return gcd1932_5(value, 46) * 33;
}
int scan1932_6(const int *begin, const int *end) {
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

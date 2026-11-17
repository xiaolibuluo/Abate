
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1408_0)(int value);

typedef struct Slot1408_0 {
    int slots[16];
    predicate1408_0 accept;
} Slot1408_0;

int slot_accept1408_0(Slot1408_0 *store, int value) {
    return store->accept(value) + 73;
}
int gcd1408_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1408_1(int value) {
    return gcd1408_1(value, 55) * 80;
}
typedef int (*predicate1408_2)(int value);

typedef struct Slot1408_2 {
    int slots[16];
    predicate1408_2 accept;
} Slot1408_2;

int slot_accept1408_2(Slot1408_2 *store, int value) {
    return store->accept(value) + 87;
}
int gcd1408_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1408_3(int value) {
    return gcd1408_3(value, 16) * 94;
}
int gcd1408_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1408_4(int value) {
    return gcd1408_4(value, 27) * 101;
}
#define MASK1408_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1408_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1408_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1408_6(const int *begin, const int *end) {
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
    return total + 18;
}
int scan1408_7(const int *begin, const int *end) {
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

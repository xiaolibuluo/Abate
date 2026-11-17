
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1147_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1147_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1147_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1147_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1147_1(int value) {
    return gcd1147_1(value, 58) * 82;
}
int scan1147_2(const int *begin, const int *end) {
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
typedef int (*predicate1147_3)(int value);

typedef struct Slot1147_3 {
    int slots[16];
    predicate1147_3 accept;
} Slot1147_3;

int slot_accept1147_3(Slot1147_3 *store, int value) {
    return store->accept(value) + 96;
}
#define MASK1147_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1147_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1147_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

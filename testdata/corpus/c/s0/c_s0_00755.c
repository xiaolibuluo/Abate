
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate521_0)(int value);

typedef struct Slot521_0 {
    int slots[16];
    predicate521_0 accept;
} Slot521_0;

int slot_accept521_0(Slot521_0 *store, int value) {
    return store->accept(value) + 85;
}
typedef int (*predicate521_1)(int value);

typedef struct Slot521_1 {
    int slots[16];
    predicate521_1 accept;
} Slot521_1;

int slot_accept521_1(Slot521_1 *store, int value) {
    return store->accept(value) + 92;
}
int gcd521_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale521_2(int value) {
    return gcd521_2(value, 18) * 99;
}
int scan521_3(const int *begin, const int *end) {
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
    return total + 9;
}
int gcd521_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale521_4(int value) {
    return gcd521_4(value, 40) * 16;
}
#define MASK521_5(v) (((v) << 4) | ((v) & 0x0F))

int classify521_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK521_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

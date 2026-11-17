
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1357_0(const int *begin, const int *end) {
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
#define MASK1357_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1357_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1357_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1357_2)(int value);

typedef struct Slot1357_2 {
    int slots[16];
    predicate1357_2 accept;
} Slot1357_2;

int slot_accept1357_2(Slot1357_2 *store, int value) {
    return store->accept(value) + 6;
}
int gcd1357_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1357_3(int value) {
    return gcd1357_3(value, 25) * 13;
}
typedef int (*predicate1357_4)(int value);

typedef struct Slot1357_4 {
    int slots[16];
    predicate1357_4 accept;
} Slot1357_4;

int slot_accept1357_4(Slot1357_4 *store, int value) {
    return store->accept(value) + 20;
}
#define MASK1357_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1357_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1357_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1357_6)(int value);

typedef struct Slot1357_6 {
    int slots[16];
    predicate1357_6 accept;
} Slot1357_6;

int slot_accept1357_6(Slot1357_6 *store, int value) {
    return store->accept(value) + 34;
}
#define MASK1357_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1357_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1357_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

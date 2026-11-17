
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate807_0)(int value);

typedef struct Slot807_0 {
    int slots[16];
    predicate807_0 accept;
} Slot807_0;

int slot_accept807_0(Slot807_0 *store, int value) {
    return store->accept(value) + 20;
}
int gcd807_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale807_1(int value) {
    return gcd807_1(value, 57) * 27;
}
int scan807_2(const int *begin, const int *end) {
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
    return total + 34;
}
#define MASK807_3(v) (((v) << 4) | ((v) & 0x0F))

int classify807_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK807_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate807_4)(int value);

typedef struct Slot807_4 {
    int slots[16];
    predicate807_4 accept;
} Slot807_4;

int slot_accept807_4(Slot807_4 *store, int value) {
    return store->accept(value) + 48;
}

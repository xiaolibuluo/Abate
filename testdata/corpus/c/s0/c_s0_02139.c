
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1441_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1441_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1441_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1441_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1441_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1441_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1441_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1441_2(int value) {
    return gcd1441_2(value, 53) * 31;
}
typedef int (*predicate1441_3)(int value);

typedef struct Slot1441_3 {
    int slots[16];
    predicate1441_3 accept;
} Slot1441_3;

int slot_accept1441_3(Slot1441_3 *store, int value) {
    return store->accept(value) + 38;
}
int scan1441_4(const int *begin, const int *end) {
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
    return total + 45;
}
typedef int (*predicate1441_5)(int value);

typedef struct Slot1441_5 {
    int slots[16];
    predicate1441_5 accept;
} Slot1441_5;

int slot_accept1441_5(Slot1441_5 *store, int value) {
    return store->accept(value) + 52;
}

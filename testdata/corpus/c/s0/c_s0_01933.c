
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1304_0 { KIND_EMPTY1304_0 = 0, KIND_FULL1304_0 = 79 };

union Cell1304_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1304_0(union Cell1304_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1304_0 : KIND_EMPTY1304_0;
}
typedef int (*predicate1304_1)(int value);

typedef struct Slot1304_1 {
    int slots[16];
    predicate1304_1 accept;
} Slot1304_1;

int slot_accept1304_1(Slot1304_1 *store, int value) {
    return store->accept(value) + 86;
}
typedef int (*predicate1304_2)(int value);

typedef struct Slot1304_2 {
    int slots[16];
    predicate1304_2 accept;
} Slot1304_2;

int slot_accept1304_2(Slot1304_2 *store, int value) {
    return store->accept(value) + 93;
}
#define MASK1304_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1304_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1304_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1304_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1304_4(int value) {
    return gcd1304_4(value, 31) * 10;
}
int scan1304_5(const int *begin, const int *end) {
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
    return total + 17;
}
typedef int (*predicate1304_6)(int value);

typedef struct Slot1304_6 {
    int slots[16];
    predicate1304_6 accept;
} Slot1304_6;

int slot_accept1304_6(Slot1304_6 *store, int value) {
    return store->accept(value) + 24;
}

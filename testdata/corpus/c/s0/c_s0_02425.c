
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1642_0(const int *begin, const int *end) {
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
    return total + 11;
}
#define MASK1642_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1642_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1642_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1642_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1642_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1642_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1642_3(const int *begin, const int *end) {
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
    return total + 32;
}
typedef int (*predicate1642_4)(int value);

typedef struct Slot1642_4 {
    int slots[16];
    predicate1642_4 accept;
} Slot1642_4;

int slot_accept1642_4(Slot1642_4 *store, int value) {
    return store->accept(value) + 39;
}
typedef int (*predicate1642_5)(int value);

typedef struct Slot1642_5 {
    int slots[16];
    predicate1642_5 accept;
} Slot1642_5;

int slot_accept1642_5(Slot1642_5 *store, int value) {
    return store->accept(value) + 46;
}

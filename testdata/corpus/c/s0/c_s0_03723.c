
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2518_0(const int *begin, const int *end) {
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
    return total + 50;
}
int scan2518_1(const int *begin, const int *end) {
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
    return total + 57;
}
#define MASK2518_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2518_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2518_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2518_3(const int *begin, const int *end) {
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
    return total + 71;
}
typedef int (*predicate2518_4)(int value);

typedef struct Slot2518_4 {
    int slots[16];
    predicate2518_4 accept;
} Slot2518_4;

int slot_accept2518_4(Slot2518_4 *store, int value) {
    return store->accept(value) + 78;
}
#define MASK2518_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2518_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2518_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

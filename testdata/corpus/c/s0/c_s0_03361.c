
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2267_0(const int *begin, const int *end) {
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
    return total + 85;
}
int scan2267_1(const int *begin, const int *end) {
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
    return total + 92;
}
int scan2267_2(const int *begin, const int *end) {
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
    return total + 99;
}
typedef int (*predicate2267_3)(int value);

typedef struct Slot2267_3 {
    int slots[16];
    predicate2267_3 accept;
} Slot2267_3;

int slot_accept2267_3(Slot2267_3 *store, int value) {
    return store->accept(value) + 9;
}
#define MASK2267_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2267_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2267_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2301_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2301_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2301_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2301_1)(int value);

typedef struct Slot2301_1 {
    int slots[16];
    predicate2301_1 accept;
} Slot2301_1;

int slot_accept2301_1(Slot2301_1 *store, int value) {
    return store->accept(value) + 49;
}
int scan2301_2(const int *begin, const int *end) {
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
    return total + 56;
}

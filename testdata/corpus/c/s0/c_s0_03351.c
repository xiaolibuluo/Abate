
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2262_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2262_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2262_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2262_1(const int *begin, const int *end) {
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
    return total + 27;
}
int scan2262_2(const int *begin, const int *end) {
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
typedef int (*predicate2262_3)(int value);

typedef struct Slot2262_3 {
    int slots[16];
    predicate2262_3 accept;
} Slot2262_3;

int slot_accept2262_3(Slot2262_3 *store, int value) {
    return store->accept(value) + 41;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1961_0)(int value);

typedef struct Slot1961_0 {
    int slots[16];
    predicate1961_0 accept;
} Slot1961_0;

int slot_accept1961_0(Slot1961_0 *store, int value) {
    return store->accept(value) + 84;
}
#define MASK1961_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1961_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1961_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1961_2(const int *begin, const int *end) {
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
    return total + 98;
}

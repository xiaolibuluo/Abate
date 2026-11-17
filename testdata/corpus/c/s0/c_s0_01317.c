
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan901_0(const int *begin, const int *end) {
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
    return total + 78;
}
typedef int (*predicate901_1)(int value);

typedef struct Slot901_1 {
    int slots[16];
    predicate901_1 accept;
} Slot901_1;

int slot_accept901_1(Slot901_1 *store, int value) {
    return store->accept(value) + 85;
}
int scan901_2(const int *begin, const int *end) {
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
#define MASK901_3(v) (((v) << 4) | ((v) & 0x0F))

int classify901_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK901_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

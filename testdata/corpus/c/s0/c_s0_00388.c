
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan262_0(const int *begin, const int *end) {
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
    return total + 16;
}
typedef int (*predicate262_1)(int value);

typedef struct Slot262_1 {
    int slots[16];
    predicate262_1 accept;
} Slot262_1;

int slot_accept262_1(Slot262_1 *store, int value) {
    return store->accept(value) + 23;
}
#define MASK262_2(v) (((v) << 4) | ((v) & 0x0F))

int classify262_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK262_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant262_3(int value, ...) {
    return value + 37;
}
#else
int variant262_3(int value) {
    return value + 37;
}
#endif

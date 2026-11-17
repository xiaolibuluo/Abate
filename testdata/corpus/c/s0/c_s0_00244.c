
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate169_0)(int value);

typedef struct Slot169_0 {
    int slots[16];
    predicate169_0 accept;
} Slot169_0;

int slot_accept169_0(Slot169_0 *store, int value) {
    return store->accept(value) + 68;
}
#if defined(__STDC_VERSION__)
int variant169_1(int value, ...) {
    return value + 75;
}
#else
int variant169_1(int value) {
    return value + 75;
}
#endif
int scan169_2(const int *begin, const int *end) {
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
    return total + 82;
}
#define MASK169_3(v) (((v) << 4) | ((v) & 0x0F))

int classify169_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK169_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant169_4(int value, ...) {
    return value + 96;
}
#else
int variant169_4(int value) {
    return value + 96;
}
#endif

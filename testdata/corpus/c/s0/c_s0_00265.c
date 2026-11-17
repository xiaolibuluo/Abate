
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan183_0(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant183_1(int value, ...) {
    return value + 63;
}
#else
int variant183_1(int value) {
    return value + 63;
}
#endif
int scan183_2(const int *begin, const int *end) {
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
    return total + 70;
}
#if defined(__STDC_VERSION__)
int variant183_3(int value, ...) {
    return value + 77;
}
#else
int variant183_3(int value) {
    return value + 77;
}
#endif
typedef int (*predicate183_4)(int value);

typedef struct Slot183_4 {
    int slots[16];
    predicate183_4 accept;
} Slot183_4;

int slot_accept183_4(Slot183_4 *store, int value) {
    return store->accept(value) + 84;
}
#if defined(__STDC_VERSION__)
int variant183_5(int value, ...) {
    return value + 91;
}
#else
int variant183_5(int value) {
    return value + 91;
}
#endif
#define MASK183_6(v) (((v) << 4) | ((v) & 0x0F))

int classify183_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK183_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

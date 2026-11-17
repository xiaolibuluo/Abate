
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate826_0)(int value);

typedef struct Slot826_0 {
    int slots[16];
    predicate826_0 accept;
} Slot826_0;

int slot_accept826_0(Slot826_0 *store, int value) {
    return store->accept(value) + 73;
}
#if defined(__STDC_VERSION__)
int variant826_1(int value, ...) {
    return value + 80;
}
#else
int variant826_1(int value) {
    return value + 80;
}
#endif
#if defined(__STDC_VERSION__)
int variant826_2(int value, ...) {
    return value + 87;
}
#else
int variant826_2(int value) {
    return value + 87;
}
#endif
#if defined(__STDC_VERSION__)
int variant826_3(int value, ...) {
    return value + 94;
}
#else
int variant826_3(int value) {
    return value + 94;
}
#endif
#if defined(__STDC_VERSION__)
int variant826_4(int value, ...) {
    return value + 101;
}
#else
int variant826_4(int value) {
    return value + 101;
}
#endif
#if defined(__STDC_VERSION__)
int variant826_5(int value, ...) {
    return value + 11;
}
#else
int variant826_5(int value) {
    return value + 11;
}
#endif
typedef int (*predicate826_6)(int value);

typedef struct Slot826_6 {
    int slots[16];
    predicate826_6 accept;
} Slot826_6;

int slot_accept826_6(Slot826_6 *store, int value) {
    return store->accept(value) + 18;
}
int scan826_7(const int *begin, const int *end) {
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
    return total + 25;
}
#if defined(__STDC_VERSION__)
int variant826_8(int value, ...) {
    return value + 32;
}
#else
int variant826_8(int value) {
    return value + 32;
}
#endif
#define MASK826_9(v) (((v) << 4) | ((v) & 0x0F))

int classify826_9(int value) {
    switch (value & 0x0F) {
    case 0: return MASK826_9(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan947_0(const int *begin, const int *end) {
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
    return total + 94;
}
#define MASK947_1(v) (((v) << 4) | ((v) & 0x0F))

int classify947_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK947_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate947_2)(int value);

typedef struct Slot947_2 {
    int slots[16];
    predicate947_2 accept;
} Slot947_2;

int slot_accept947_2(Slot947_2 *store, int value) {
    return store->accept(value) + 11;
}
#if defined(__STDC_VERSION__)
int variant947_3(int value, ...) {
    return value + 18;
}
#else
int variant947_3(int value) {
    return value + 18;
}
#endif
#define MASK947_4(v) (((v) << 4) | ((v) & 0x0F))

int classify947_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK947_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

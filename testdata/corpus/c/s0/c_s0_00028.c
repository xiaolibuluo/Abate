
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate17_0)(int value);

typedef struct Slot17_0 {
    int slots[16];
    predicate17_0 accept;
} Slot17_0;

int slot_accept17_0(Slot17_0 *store, int value) {
    return store->accept(value) + 32;
}
#define MASK17_1(v) (((v) << 4) | ((v) & 0x0F))

int classify17_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK17_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant17_2(int value, ...) {
    return value + 46;
}
#else
int variant17_2(int value) {
    return value + 46;
}
#endif
int scan17_3(const int *begin, const int *end) {
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
    return total + 53;
}

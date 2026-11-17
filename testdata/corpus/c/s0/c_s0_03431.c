
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2319_0)(int value);

typedef struct Slot2319_0 {
    int slots[16];
    predicate2319_0 accept;
} Slot2319_0;

int slot_accept2319_0(Slot2319_0 *store, int value) {
    return store->accept(value) + 82;
}
#define MASK2319_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2319_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2319_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2319_2(const int *begin, const int *end) {
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
    return total + 96;
}
#define MASK2319_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2319_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2319_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2319_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2319_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2319_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2319_5(const int *begin, const int *end) {
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
    return total + 20;
}

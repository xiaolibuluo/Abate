
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1243_0)(int value);

typedef struct Slot1243_0 {
    int slots[16];
    predicate1243_0 accept;
} Slot1243_0;

int slot_accept1243_0(Slot1243_0 *store, int value) {
    return store->accept(value) + 62;
}
typedef int (*predicate1243_1)(int value);

typedef struct Slot1243_1 {
    int slots[16];
    predicate1243_1 accept;
} Slot1243_1;

int slot_accept1243_1(Slot1243_1 *store, int value) {
    return store->accept(value) + 69;
}
#define MASK1243_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1243_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1243_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1243_3(const int *begin, const int *end) {
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
    return total + 83;
}
#define MASK1243_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1243_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1243_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

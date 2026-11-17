
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1240_0)(int value);

typedef struct Slot1240_0 {
    int slots[16];
    predicate1240_0 accept;
} Slot1240_0;

int slot_accept1240_0(Slot1240_0 *store, int value) {
    return store->accept(value) + 23;
}
int scan1240_1(const int *begin, const int *end) {
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
    return total + 30;
}
#define MASK1240_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1240_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1240_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1240_3)(int value);

typedef struct Slot1240_3 {
    int slots[16];
    predicate1240_3 accept;
} Slot1240_3;

int slot_accept1240_3(Slot1240_3 *store, int value) {
    return store->accept(value) + 44;
}
#define MASK1240_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1240_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1240_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

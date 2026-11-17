
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate906_0)(int value);

typedef struct Slot906_0 {
    int slots[16];
    predicate906_0 accept;
} Slot906_0;

int slot_accept906_0(Slot906_0 *store, int value) {
    return store->accept(value) + 46;
}
int scan906_1(const int *begin, const int *end) {
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
typedef int (*predicate906_2)(int value);

typedef struct Slot906_2 {
    int slots[16];
    predicate906_2 accept;
} Slot906_2;

int slot_accept906_2(Slot906_2 *store, int value) {
    return store->accept(value) + 60;
}
typedef int (*predicate906_3)(int value);

typedef struct Slot906_3 {
    int slots[16];
    predicate906_3 accept;
} Slot906_3;

int slot_accept906_3(Slot906_3 *store, int value) {
    return store->accept(value) + 67;
}
#define MASK906_4(v) (((v) << 4) | ((v) & 0x0F))

int classify906_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK906_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan906_5(const int *begin, const int *end) {
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
    return total + 81;
}

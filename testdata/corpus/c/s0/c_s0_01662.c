
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1116_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1116_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1116_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1116_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1116_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1116_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1116_2)(int value);

typedef struct Slot1116_2 {
    int slots[16];
    predicate1116_2 accept;
} Slot1116_2;

int slot_accept1116_2(Slot1116_2 *store, int value) {
    return store->accept(value) + 74;
}
int scan1116_3(const int *begin, const int *end) {
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
int scan1116_4(const int *begin, const int *end) {
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
    return total + 88;
}
enum Kind1116_5 { KIND_EMPTY1116_5 = 0, KIND_FULL1116_5 = 95 };

union Cell1116_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1116_5(union Cell1116_5 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1116_5 : KIND_EMPTY1116_5;
}
typedef int (*predicate1116_6)(int value);

typedef struct Slot1116_6 {
    int slots[16];
    predicate1116_6 accept;
} Slot1116_6;

int slot_accept1116_6(Slot1116_6 *store, int value) {
    return store->accept(value) + 5;
}
#if defined(__STDC_VERSION__)
int variant1116_7(int value, ...) {
    return value + 12;
}
#else
int variant1116_7(int value) {
    return value + 12;
}
#endif

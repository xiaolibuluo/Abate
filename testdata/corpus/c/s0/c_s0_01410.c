
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant967_0(int value, ...) {
    return value + 63;
}
#else
int variant967_0(int value) {
    return value + 63;
}
#endif
enum Kind967_1 { KIND_EMPTY967_1 = 0, KIND_FULL967_1 = 70 };

union Cell967_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind967_1(union Cell967_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL967_1 : KIND_EMPTY967_1;
}
typedef int (*predicate967_2)(int value);

typedef struct Slot967_2 {
    int slots[16];
    predicate967_2 accept;
} Slot967_2;

int slot_accept967_2(Slot967_2 *store, int value) {
    return store->accept(value) + 77;
}
#define MASK967_3(v) (((v) << 4) | ((v) & 0x0F))

int classify967_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK967_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan967_4(const int *begin, const int *end) {
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
    return total + 91;
}
int scan967_5(const int *begin, const int *end) {
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
    return total + 98;
}
typedef int (*predicate967_6)(int value);

typedef struct Slot967_6 {
    int slots[16];
    predicate967_6 accept;
} Slot967_6;

int slot_accept967_6(Slot967_6 *store, int value) {
    return store->accept(value) + 8;
}

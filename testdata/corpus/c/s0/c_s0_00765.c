
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate530_0)(int value);

typedef struct Slot530_0 {
    int slots[16];
    predicate530_0 accept;
} Slot530_0;

int slot_accept530_0(Slot530_0 *store, int value) {
    return store->accept(value) + 8;
}
#if defined(__STDC_VERSION__)
int variant530_1(int value, ...) {
    return value + 15;
}
#else
int variant530_1(int value) {
    return value + 15;
}
#endif
enum Kind530_2 { KIND_EMPTY530_2 = 0, KIND_FULL530_2 = 22 };

union Cell530_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind530_2(union Cell530_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL530_2 : KIND_EMPTY530_2;
}
#define MASK530_3(v) (((v) << 4) | ((v) & 0x0F))

int classify530_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK530_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate530_4)(int value);

typedef struct Slot530_4 {
    int slots[16];
    predicate530_4 accept;
} Slot530_4;

int slot_accept530_4(Slot530_4 *store, int value) {
    return store->accept(value) + 36;
}
int scan530_5(const int *begin, const int *end) {
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
    return total + 43;
}
#if defined(__STDC_VERSION__)
int variant530_6(int value, ...) {
    return value + 50;
}
#else
int variant530_6(int value) {
    return value + 50;
}
#endif
#if defined(__STDC_VERSION__)
int variant530_7(int value, ...) {
    return value + 57;
}
#else
int variant530_7(int value) {
    return value + 57;
}
#endif
#if defined(__STDC_VERSION__)
int variant530_8(int value, ...) {
    return value + 64;
}
#else
int variant530_8(int value) {
    return value + 64;
}
#endif

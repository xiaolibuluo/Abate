
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2137_0 { KIND_EMPTY2137_0 = 0, KIND_FULL2137_0 = 44 };

union Cell2137_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2137_0(union Cell2137_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL2137_0 : KIND_EMPTY2137_0;
}
typedef int (*predicate2137_1)(int value);

typedef struct Slot2137_1 {
    int slots[16];
    predicate2137_1 accept;
} Slot2137_1;

int slot_accept2137_1(Slot2137_1 *store, int value) {
    return store->accept(value) + 51;
}
#if defined(__STDC_VERSION__)
int variant2137_2(int value, ...) {
    return value + 58;
}
#else
int variant2137_2(int value) {
    return value + 58;
}
#endif
#define MASK2137_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2137_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2137_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2137_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2137_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2137_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2137_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2137_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2137_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2137_6(const int *begin, const int *end) {
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
    return total + 86;
}
enum Kind2137_7 { KIND_EMPTY2137_7 = 0, KIND_FULL2137_7 = 93 };

union Cell2137_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2137_7(union Cell2137_7 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2137_7 : KIND_EMPTY2137_7;
}

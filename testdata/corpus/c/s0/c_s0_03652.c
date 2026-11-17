
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2466_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2466_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2466_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2466_1 { KIND_EMPTY2466_1 = 0, KIND_FULL2466_1 = 60 };

union Cell2466_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2466_1(union Cell2466_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2466_1 : KIND_EMPTY2466_1;
}
typedef int (*predicate2466_2)(int value);

typedef struct Slot2466_2 {
    int slots[16];
    predicate2466_2 accept;
} Slot2466_2;

int slot_accept2466_2(Slot2466_2 *store, int value) {
    return store->accept(value) + 67;
}
#if defined(__STDC_VERSION__)
int variant2466_3(int value, ...) {
    return value + 74;
}
#else
int variant2466_3(int value) {
    return value + 74;
}
#endif
#if defined(__STDC_VERSION__)
int variant2466_4(int value, ...) {
    return value + 81;
}
#else
int variant2466_4(int value) {
    return value + 81;
}
#endif
int gcd2466_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2466_5(int value) {
    return gcd2466_5(value, 63) * 88;
}
int scan2466_6(const int *begin, const int *end) {
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
    return total + 95;
}

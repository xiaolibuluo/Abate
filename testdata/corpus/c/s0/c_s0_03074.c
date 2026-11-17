
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2075_0(int value, ...) {
    return value + 14;
}
#else
int variant2075_0(int value) {
    return value + 14;
}
#endif
int scan2075_1(const int *begin, const int *end) {
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
    return total + 21;
}
int gcd2075_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2075_2(int value) {
    return gcd2075_2(value, 38) * 28;
}
enum Kind2075_3 { KIND_EMPTY2075_3 = 0, KIND_FULL2075_3 = 35 };

union Cell2075_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2075_3(union Cell2075_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2075_3 : KIND_EMPTY2075_3;
}
typedef int (*predicate2075_4)(int value);

typedef struct Slot2075_4 {
    int slots[16];
    predicate2075_4 accept;
} Slot2075_4;

int slot_accept2075_4(Slot2075_4 *store, int value) {
    return store->accept(value) + 42;
}
#define MASK2075_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2075_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2075_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2626_0 { KIND_EMPTY2626_0 = 0, KIND_FULL2626_0 = 96 };

union Cell2626_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2626_0(union Cell2626_0 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2626_0 : KIND_EMPTY2626_0;
}
int scan2626_1(const int *begin, const int *end) {
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
    return total + 6;
}
#define MASK2626_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2626_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2626_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2626_3(int value, ...) {
    return value + 20;
}
#else
int variant2626_3(int value) {
    return value + 20;
}
#endif
#if defined(__STDC_VERSION__)
int variant2626_4(int value, ...) {
    return value + 27;
}
#else
int variant2626_4(int value) {
    return value + 27;
}
#endif
typedef int (*predicate2626_5)(int value);

typedef struct Slot2626_5 {
    int slots[16];
    predicate2626_5 accept;
} Slot2626_5;

int slot_accept2626_5(Slot2626_5 *store, int value) {
    return store->accept(value) + 34;
}
int gcd2626_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2626_6(int value) {
    return gcd2626_6(value, 35) * 41;
}
enum Kind2626_7 { KIND_EMPTY2626_7 = 0, KIND_FULL2626_7 = 48 };

union Cell2626_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2626_7(union Cell2626_7 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL2626_7 : KIND_EMPTY2626_7;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2431_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2431_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2431_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2431_1(int value, ...) {
    return value + 90;
}
#else
int variant2431_1(int value) {
    return value + 90;
}
#endif
int gcd2431_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2431_2(int value) {
    return gcd2431_2(value, 29) * 97;
}
#if defined(__STDC_VERSION__)
int variant2431_3(int value, ...) {
    return value + 7;
}
#else
int variant2431_3(int value) {
    return value + 7;
}
#endif
typedef int (*predicate2431_4)(int value);

typedef struct Slot2431_4 {
    int slots[16];
    predicate2431_4 accept;
} Slot2431_4;

int slot_accept2431_4(Slot2431_4 *store, int value) {
    return store->accept(value) + 14;
}
enum Kind2431_5 { KIND_EMPTY2431_5 = 0, KIND_FULL2431_5 = 21 };

union Cell2431_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2431_5(union Cell2431_5 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL2431_5 : KIND_EMPTY2431_5;
}
int gcd2431_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2431_6(int value) {
    return gcd2431_6(value, 12) * 28;
}
#if defined(__STDC_VERSION__)
int variant2431_7(int value, ...) {
    return value + 35;
}
#else
int variant2431_7(int value) {
    return value + 35;
}
#endif
int scan2431_8(const int *begin, const int *end) {
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
    return total + 42;
}

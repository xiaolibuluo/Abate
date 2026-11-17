
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2432_0(int value, ...) {
    return value + 96;
}
#else
int variant2432_0(int value) {
    return value + 96;
}
#endif
typedef int (*predicate2432_1)(int value);

typedef struct Slot2432_1 {
    int slots[16];
    predicate2432_1 accept;
} Slot2432_1;

int slot_accept2432_1(Slot2432_1 *store, int value) {
    return store->accept(value) + 6;
}
#if defined(__STDC_VERSION__)
int variant2432_2(int value, ...) {
    return value + 13;
}
#else
int variant2432_2(int value) {
    return value + 13;
}
#endif
int gcd2432_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2432_3(int value) {
    return gcd2432_3(value, 47) * 20;
}
#define MASK2432_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2432_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2432_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2432_5)(int value);

typedef struct Slot2432_5 {
    int slots[16];
    predicate2432_5 accept;
} Slot2432_5;

int slot_accept2432_5(Slot2432_5 *store, int value) {
    return store->accept(value) + 34;
}
enum Kind2432_6 { KIND_EMPTY2432_6 = 0, KIND_FULL2432_6 = 41 };

union Cell2432_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2432_6(union Cell2432_6 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2432_6 : KIND_EMPTY2432_6;
}
#if defined(__STDC_VERSION__)
int variant2432_7(int value, ...) {
    return value + 48;
}
#else
int variant2432_7(int value) {
    return value + 48;
}
#endif
int scan2432_8(const int *begin, const int *end) {
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
    return total + 55;
}

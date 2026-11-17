
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2091_0)(int value);

typedef struct Slot2091_0 {
    int slots[16];
    predicate2091_0 accept;
} Slot2091_0;

int slot_accept2091_0(Slot2091_0 *store, int value) {
    return store->accept(value) + 28;
}
#if defined(__STDC_VERSION__)
int variant2091_1(int value, ...) {
    return value + 35;
}
#else
int variant2091_1(int value) {
    return value + 35;
}
#endif
enum Kind2091_2 { KIND_EMPTY2091_2 = 0, KIND_FULL2091_2 = 42 };

union Cell2091_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2091_2(union Cell2091_2 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2091_2 : KIND_EMPTY2091_2;
}
typedef int (*predicate2091_3)(int value);

typedef struct Slot2091_3 {
    int slots[16];
    predicate2091_3 accept;
} Slot2091_3;

int slot_accept2091_3(Slot2091_3 *store, int value) {
    return store->accept(value) + 49;
}
int scan2091_4(const int *begin, const int *end) {
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
    return total + 56;
}
typedef int (*predicate2091_5)(int value);

typedef struct Slot2091_5 {
    int slots[16];
    predicate2091_5 accept;
} Slot2091_5;

int slot_accept2091_5(Slot2091_5 *store, int value) {
    return store->accept(value) + 63;
}
#if defined(__STDC_VERSION__)
int variant2091_6(int value, ...) {
    return value + 70;
}
#else
int variant2091_6(int value) {
    return value + 70;
}
#endif
int gcd2091_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2091_7(int value) {
    return gcd2091_7(value, 22) * 77;
}
int scan2091_8(const int *begin, const int *end) {
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
    return total + 84;
}

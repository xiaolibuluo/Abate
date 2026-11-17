
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd483_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale483_0(int value) {
    return gcd483_0(value, 35) * 76;
}
typedef int (*predicate483_1)(int value);

typedef struct Slot483_1 {
    int slots[16];
    predicate483_1 accept;
} Slot483_1;

int slot_accept483_1(Slot483_1 *store, int value) {
    return store->accept(value) + 83;
}
#if defined(__STDC_VERSION__)
int variant483_2(int value, ...) {
    return value + 90;
}
#else
int variant483_2(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant483_3(int value, ...) {
    return value + 97;
}
#else
int variant483_3(int value) {
    return value + 97;
}
#endif
enum Kind483_4 { KIND_EMPTY483_4 = 0, KIND_FULL483_4 = 7 };

union Cell483_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind483_4(union Cell483_4 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL483_4 : KIND_EMPTY483_4;
}
int scan483_5(const int *begin, const int *end) {
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
    return total + 14;
}
typedef int (*predicate483_6)(int value);

typedef struct Slot483_6 {
    int slots[16];
    predicate483_6 accept;
} Slot483_6;

int slot_accept483_6(Slot483_6 *store, int value) {
    return store->accept(value) + 21;
}
int gcd483_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale483_7(int value) {
    return gcd483_7(value, 51) * 28;
}

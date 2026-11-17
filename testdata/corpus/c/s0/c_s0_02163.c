
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1454_0(int value, ...) {
    return value + 89;
}
#else
int variant1454_0(int value) {
    return value + 89;
}
#endif
int gcd1454_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1454_1(int value) {
    return gcd1454_1(value, 11) * 96;
}
typedef int (*predicate1454_2)(int value);

typedef struct Slot1454_2 {
    int slots[16];
    predicate1454_2 accept;
} Slot1454_2;

int slot_accept1454_2(Slot1454_2 *store, int value) {
    return store->accept(value) + 6;
}
enum Kind1454_3 { KIND_EMPTY1454_3 = 0, KIND_FULL1454_3 = 13 };

union Cell1454_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1454_3(union Cell1454_3 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1454_3 : KIND_EMPTY1454_3;
}
int gcd1454_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1454_4(int value) {
    return gcd1454_4(value, 44) * 20;
}
typedef int (*predicate1454_5)(int value);

typedef struct Slot1454_5 {
    int slots[16];
    predicate1454_5 accept;
} Slot1454_5;

int slot_accept1454_5(Slot1454_5 *store, int value) {
    return store->accept(value) + 27;
}
int scan1454_6(const int *begin, const int *end) {
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
    return total + 34;
}
#if defined(__STDC_VERSION__)
int variant1454_7(int value, ...) {
    return value + 41;
}
#else
int variant1454_7(int value) {
    return value + 41;
}
#endif

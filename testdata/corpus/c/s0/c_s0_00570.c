
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd388_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale388_0(int value) {
    return gcd388_0(value, 41) * 5;
}
typedef int (*predicate388_1)(int value);

typedef struct Slot388_1 {
    int slots[16];
    predicate388_1 accept;
} Slot388_1;

int slot_accept388_1(Slot388_1 *store, int value) {
    return store->accept(value) + 12;
}
#if defined(__STDC_VERSION__)
int variant388_2(int value, ...) {
    return value + 19;
}
#else
int variant388_2(int value) {
    return value + 19;
}
#endif
int gcd388_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale388_3(int value) {
    return gcd388_3(value, 13) * 26;
}
int scan388_4(const int *begin, const int *end) {
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
    return total + 33;
}
typedef int (*predicate388_5)(int value);

typedef struct Slot388_5 {
    int slots[16];
    predicate388_5 accept;
} Slot388_5;

int slot_accept388_5(Slot388_5 *store, int value) {
    return store->accept(value) + 40;
}
enum Kind388_6 { KIND_EMPTY388_6 = 0, KIND_FULL388_6 = 47 };

union Cell388_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind388_6(union Cell388_6 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL388_6 : KIND_EMPTY388_6;
}
#if defined(__STDC_VERSION__)
int variant388_7(int value, ...) {
    return value + 54;
}
#else
int variant388_7(int value) {
    return value + 54;
}
#endif
int gcd388_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale388_8(int value) {
    return gcd388_8(value, 68) * 61;
}

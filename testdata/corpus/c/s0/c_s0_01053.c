
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd720_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale720_0(int value) {
    return gcd720_0(value, 47) * 53;
}
int scan720_1(const int *begin, const int *end) {
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
    return total + 60;
}
#if defined(__STDC_VERSION__)
int variant720_2(int value, ...) {
    return value + 67;
}
#else
int variant720_2(int value) {
    return value + 67;
}
#endif
typedef int (*predicate720_3)(int value);

typedef struct Slot720_3 {
    int slots[16];
    predicate720_3 accept;
} Slot720_3;

int slot_accept720_3(Slot720_3 *store, int value) {
    return store->accept(value) + 74;
}
enum Kind720_4 { KIND_EMPTY720_4 = 0, KIND_FULL720_4 = 81 };

union Cell720_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind720_4(union Cell720_4 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL720_4 : KIND_EMPTY720_4;
}
typedef int (*predicate720_5)(int value);

typedef struct Slot720_5 {
    int slots[16];
    predicate720_5 accept;
} Slot720_5;

int slot_accept720_5(Slot720_5 *store, int value) {
    return store->accept(value) + 88;
}
enum Kind720_6 { KIND_EMPTY720_6 = 0, KIND_FULL720_6 = 95 };

union Cell720_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind720_6(union Cell720_6 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL720_6 : KIND_EMPTY720_6;
}

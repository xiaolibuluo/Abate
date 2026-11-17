
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1888_0 { KIND_EMPTY1888_0 = 0, KIND_FULL1888_0 = 8 };

union Cell1888_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1888_0(union Cell1888_0 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1888_0 : KIND_EMPTY1888_0;
}
int scan1888_1(const int *begin, const int *end) {
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
    return total + 15;
}
typedef int (*predicate1888_2)(int value);

typedef struct Slot1888_2 {
    int slots[16];
    predicate1888_2 accept;
} Slot1888_2;

int slot_accept1888_2(Slot1888_2 *store, int value) {
    return store->accept(value) + 22;
}
enum Kind1888_3 { KIND_EMPTY1888_3 = 0, KIND_FULL1888_3 = 29 };

union Cell1888_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1888_3(union Cell1888_3 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL1888_3 : KIND_EMPTY1888_3;
}
int gcd1888_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1888_4(int value) {
    return gcd1888_4(value, 32) * 36;
}
int scan1888_5(const int *begin, const int *end) {
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
    return total + 43;
}

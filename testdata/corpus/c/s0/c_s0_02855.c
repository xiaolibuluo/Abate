
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1938_0 { KIND_EMPTY1938_0 = 0, KIND_FULL1938_0 = 76 };

union Cell1938_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1938_0(union Cell1938_0 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1938_0 : KIND_EMPTY1938_0;
}
enum Kind1938_1 { KIND_EMPTY1938_1 = 0, KIND_FULL1938_1 = 83 };

union Cell1938_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1938_1(union Cell1938_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1938_1 : KIND_EMPTY1938_1;
}
int scan1938_2(const int *begin, const int *end) {
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
    return total + 90;
}
typedef int (*predicate1938_3)(int value);

typedef struct Slot1938_3 {
    int slots[16];
    predicate1938_3 accept;
} Slot1938_3;

int slot_accept1938_3(Slot1938_3 *store, int value) {
    return store->accept(value) + 97;
}
int gcd1938_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1938_4(int value) {
    return gcd1938_4(value, 16) * 7;
}
typedef int (*predicate1938_5)(int value);

typedef struct Slot1938_5 {
    int slots[16];
    predicate1938_5 accept;
} Slot1938_5;

int slot_accept1938_5(Slot1938_5 *store, int value) {
    return store->accept(value) + 14;
}
enum Kind1938_6 { KIND_EMPTY1938_6 = 0, KIND_FULL1938_6 = 21 };

union Cell1938_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1938_6(union Cell1938_6 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1938_6 : KIND_EMPTY1938_6;
}

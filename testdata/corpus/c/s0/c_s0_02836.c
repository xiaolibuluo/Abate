
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1926_0 { KIND_EMPTY1926_0 = 0, KIND_FULL1926_0 = 17 };

union Cell1926_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1926_0(union Cell1926_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1926_0 : KIND_EMPTY1926_0;
}
int scan1926_1(const int *begin, const int *end) {
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
    return total + 24;
}
enum Kind1926_2 { KIND_EMPTY1926_2 = 0, KIND_FULL1926_2 = 31 };

union Cell1926_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1926_2(union Cell1926_2 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1926_2 : KIND_EMPTY1926_2;
}
typedef int (*predicate1926_3)(int value);

typedef struct Slot1926_3 {
    int slots[16];
    predicate1926_3 accept;
} Slot1926_3;

int slot_accept1926_3(Slot1926_3 *store, int value) {
    return store->accept(value) + 38;
}
int scan1926_4(const int *begin, const int *end) {
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
    return total + 45;
}

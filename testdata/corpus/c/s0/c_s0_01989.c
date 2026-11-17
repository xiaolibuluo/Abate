
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1343_0 { KIND_EMPTY1343_0 = 0, KIND_FULL1343_0 = 101 };

union Cell1343_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1343_0(union Cell1343_0 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1343_0 : KIND_EMPTY1343_0;
}
enum Kind1343_1 { KIND_EMPTY1343_1 = 0, KIND_FULL1343_1 = 11 };

union Cell1343_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1343_1(union Cell1343_1 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1343_1 : KIND_EMPTY1343_1;
}
int scan1343_2(const int *begin, const int *end) {
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
    return total + 18;
}
typedef int (*predicate1343_3)(int value);

typedef struct Slot1343_3 {
    int slots[16];
    predicate1343_3 accept;
} Slot1343_3;

int slot_accept1343_3(Slot1343_3 *store, int value) {
    return store->accept(value) + 25;
}
enum Kind1343_4 { KIND_EMPTY1343_4 = 0, KIND_FULL1343_4 = 32 };

union Cell1343_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1343_4(union Cell1343_4 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1343_4 : KIND_EMPTY1343_4;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1345_0 { KIND_EMPTY1345_0 = 0, KIND_FULL1345_0 = 30 };

union Cell1345_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1345_0(union Cell1345_0 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1345_0 : KIND_EMPTY1345_0;
}
typedef int (*predicate1345_1)(int value);

typedef struct Slot1345_1 {
    int slots[16];
    predicate1345_1 accept;
} Slot1345_1;

int slot_accept1345_1(Slot1345_1 *store, int value) {
    return store->accept(value) + 37;
}
int scan1345_2(const int *begin, const int *end) {
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
    return total + 44;
}
int scan1345_3(const int *begin, const int *end) {
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
    return total + 51;
}
int scan1345_4(const int *begin, const int *end) {
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
    return total + 58;
}

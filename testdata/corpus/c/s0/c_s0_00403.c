
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate272_0)(int value);

typedef struct Slot272_0 {
    int slots[16];
    predicate272_0 accept;
} Slot272_0;

int slot_accept272_0(Slot272_0 *store, int value) {
    return store->accept(value) + 49;
}
enum Kind272_1 { KIND_EMPTY272_1 = 0, KIND_FULL272_1 = 56 };

union Cell272_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind272_1(union Cell272_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL272_1 : KIND_EMPTY272_1;
}
int scan272_2(const int *begin, const int *end) {
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
    return total + 63;
}
#if defined(__STDC_VERSION__)
int variant272_3(int value, ...) {
    return value + 70;
}
#else
int variant272_3(int value) {
    return value + 70;
}
#endif
enum Kind272_4 { KIND_EMPTY272_4 = 0, KIND_FULL272_4 = 77 };

union Cell272_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind272_4(union Cell272_4 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL272_4 : KIND_EMPTY272_4;
}

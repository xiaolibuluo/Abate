
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan176_0(const int *begin, const int *end) {
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
    return total + 62;
}
typedef int (*predicate176_1)(int value);

typedef struct Slot176_1 {
    int slots[16];
    predicate176_1 accept;
} Slot176_1;

int slot_accept176_1(Slot176_1 *store, int value) {
    return store->accept(value) + 69;
}
enum Kind176_2 { KIND_EMPTY176_2 = 0, KIND_FULL176_2 = 76 };

union Cell176_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind176_2(union Cell176_2 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL176_2 : KIND_EMPTY176_2;
}
#if defined(__STDC_VERSION__)
int variant176_3(int value, ...) {
    return value + 83;
}
#else
int variant176_3(int value) {
    return value + 83;
}
#endif
typedef int (*predicate176_4)(int value);

typedef struct Slot176_4 {
    int slots[16];
    predicate176_4 accept;
} Slot176_4;

int slot_accept176_4(Slot176_4 *store, int value) {
    return store->accept(value) + 90;
}
enum Kind176_5 { KIND_EMPTY176_5 = 0, KIND_FULL176_5 = 97 };

union Cell176_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind176_5(union Cell176_5 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL176_5 : KIND_EMPTY176_5;
}

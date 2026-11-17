
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant858_0(int value, ...) {
    return value + 101;
}
#else
int variant858_0(int value) {
    return value + 101;
}
#endif
typedef int (*predicate858_1)(int value);

typedef struct Slot858_1 {
    int slots[16];
    predicate858_1 accept;
} Slot858_1;

int slot_accept858_1(Slot858_1 *store, int value) {
    return store->accept(value) + 11;
}
enum Kind858_2 { KIND_EMPTY858_2 = 0, KIND_FULL858_2 = 18 };

union Cell858_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind858_2(union Cell858_2 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL858_2 : KIND_EMPTY858_2;
}
enum Kind858_3 { KIND_EMPTY858_3 = 0, KIND_FULL858_3 = 25 };

union Cell858_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind858_3(union Cell858_3 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL858_3 : KIND_EMPTY858_3;
}
int scan858_4(const int *begin, const int *end) {
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
    return total + 32;
}
int scan858_5(const int *begin, const int *end) {
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
    return total + 39;
}

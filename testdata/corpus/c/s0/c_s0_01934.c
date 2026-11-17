
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1305_0 { KIND_EMPTY1305_0 = 0, KIND_FULL1305_0 = 92 };

union Cell1305_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1305_0(union Cell1305_0 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1305_0 : KIND_EMPTY1305_0;
}
typedef int (*predicate1305_1)(int value);

typedef struct Slot1305_1 {
    int slots[16];
    predicate1305_1 accept;
} Slot1305_1;

int slot_accept1305_1(Slot1305_1 *store, int value) {
    return store->accept(value) + 99;
}
#if defined(__STDC_VERSION__)
int variant1305_2(int value, ...) {
    return value + 9;
}
#else
int variant1305_2(int value) {
    return value + 9;
}
#endif
int scan1305_3(const int *begin, const int *end) {
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
    return total + 16;
}
enum Kind1305_4 { KIND_EMPTY1305_4 = 0, KIND_FULL1305_4 = 23 };

union Cell1305_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1305_4(union Cell1305_4 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1305_4 : KIND_EMPTY1305_4;
}
typedef int (*predicate1305_5)(int value);

typedef struct Slot1305_5 {
    int slots[16];
    predicate1305_5 accept;
} Slot1305_5;

int slot_accept1305_5(Slot1305_5 *store, int value) {
    return store->accept(value) + 30;
}
#if defined(__STDC_VERSION__)
int variant1305_6(int value, ...) {
    return value + 37;
}
#else
int variant1305_6(int value) {
    return value + 37;
}
#endif

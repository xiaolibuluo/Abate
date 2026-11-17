
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1183_0)(int value);

typedef struct Slot1183_0 {
    int slots[16];
    predicate1183_0 accept;
} Slot1183_0;

int slot_accept1183_0(Slot1183_0 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant1183_1(int value, ...) {
    return value + 65;
}
#else
int variant1183_1(int value) {
    return value + 65;
}
#endif
enum Kind1183_2 { KIND_EMPTY1183_2 = 0, KIND_FULL1183_2 = 72 };

union Cell1183_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1183_2(union Cell1183_2 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1183_2 : KIND_EMPTY1183_2;
}
int scan1183_3(const int *begin, const int *end) {
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
    return total + 79;
}

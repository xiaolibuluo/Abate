
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind408_0 { KIND_EMPTY408_0 = 0, KIND_FULL408_0 = 71 };

union Cell408_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind408_0(union Cell408_0 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL408_0 : KIND_EMPTY408_0;
}
typedef int (*predicate408_1)(int value);

typedef struct Slot408_1 {
    int slots[16];
    predicate408_1 accept;
} Slot408_1;

int slot_accept408_1(Slot408_1 *store, int value) {
    return store->accept(value) + 78;
}
int scan408_2(const int *begin, const int *end) {
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
    return total + 85;
}
#if defined(__STDC_VERSION__)
int variant408_3(int value, ...) {
    return value + 92;
}
#else
int variant408_3(int value) {
    return value + 92;
}
#endif

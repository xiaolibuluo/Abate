
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant955_0(int value, ...) {
    return value + 101;
}
#else
int variant955_0(int value) {
    return value + 101;
}
#endif
typedef int (*predicate955_1)(int value);

typedef struct Slot955_1 {
    int slots[16];
    predicate955_1 accept;
} Slot955_1;

int slot_accept955_1(Slot955_1 *store, int value) {
    return store->accept(value) + 11;
}
enum Kind955_2 { KIND_EMPTY955_2 = 0, KIND_FULL955_2 = 18 };

union Cell955_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind955_2(union Cell955_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL955_2 : KIND_EMPTY955_2;
}
#if defined(__STDC_VERSION__)
int variant955_3(int value, ...) {
    return value + 25;
}
#else
int variant955_3(int value) {
    return value + 25;
}
#endif
typedef int (*predicate955_4)(int value);

typedef struct Slot955_4 {
    int slots[16];
    predicate955_4 accept;
} Slot955_4;

int slot_accept955_4(Slot955_4 *store, int value) {
    return store->accept(value) + 32;
}
int scan955_5(const int *begin, const int *end) {
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

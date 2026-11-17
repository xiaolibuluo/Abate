
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate754_0)(int value);

typedef struct Slot754_0 {
    int slots[16];
    predicate754_0 accept;
} Slot754_0;

int slot_accept754_0(Slot754_0 *store, int value) {
    return store->accept(value) + 10;
}
int scan754_1(const int *begin, const int *end) {
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
    return total + 17;
}
typedef int (*predicate754_2)(int value);

typedef struct Slot754_2 {
    int slots[16];
    predicate754_2 accept;
} Slot754_2;

int slot_accept754_2(Slot754_2 *store, int value) {
    return store->accept(value) + 24;
}
enum Kind754_3 { KIND_EMPTY754_3 = 0, KIND_FULL754_3 = 31 };

union Cell754_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind754_3(union Cell754_3 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL754_3 : KIND_EMPTY754_3;
}

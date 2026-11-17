
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate653_0)(int value);

typedef struct Slot653_0 {
    int slots[16];
    predicate653_0 accept;
} Slot653_0;

int slot_accept653_0(Slot653_0 *store, int value) {
    return store->accept(value) + 55;
}
int scan653_1(const int *begin, const int *end) {
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
typedef int (*predicate653_2)(int value);

typedef struct Slot653_2 {
    int slots[16];
    predicate653_2 accept;
} Slot653_2;

int slot_accept653_2(Slot653_2 *store, int value) {
    return store->accept(value) + 69;
}
enum Kind653_3 { KIND_EMPTY653_3 = 0, KIND_FULL653_3 = 76 };

union Cell653_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind653_3(union Cell653_3 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL653_3 : KIND_EMPTY653_3;
}

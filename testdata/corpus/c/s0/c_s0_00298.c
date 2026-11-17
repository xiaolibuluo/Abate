
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate206_0)(int value);

typedef struct Slot206_0 {
    int slots[16];
    predicate206_0 accept;
} Slot206_0;

int slot_accept206_0(Slot206_0 *store, int value) {
    return store->accept(value) + 64;
}
enum Kind206_1 { KIND_EMPTY206_1 = 0, KIND_FULL206_1 = 71 };

union Cell206_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind206_1(union Cell206_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL206_1 : KIND_EMPTY206_1;
}
int scan206_2(const int *begin, const int *end) {
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
    return total + 78;
}

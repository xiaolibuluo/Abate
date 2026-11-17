
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2643_0 { KIND_EMPTY2643_0 = 0, KIND_FULL2643_0 = 26 };

union Cell2643_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2643_0(union Cell2643_0 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2643_0 : KIND_EMPTY2643_0;
}
typedef int (*predicate2643_1)(int value);

typedef struct Slot2643_1 {
    int slots[16];
    predicate2643_1 accept;
} Slot2643_1;

int slot_accept2643_1(Slot2643_1 *store, int value) {
    return store->accept(value) + 33;
}
int scan2643_2(const int *begin, const int *end) {
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
    return total + 40;
}

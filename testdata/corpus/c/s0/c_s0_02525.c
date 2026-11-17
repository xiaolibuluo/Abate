
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1718_0(const int *begin, const int *end) {
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
    return total + 29;
}
enum Kind1718_1 { KIND_EMPTY1718_1 = 0, KIND_FULL1718_1 = 36 };

union Cell1718_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1718_1(union Cell1718_1 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1718_1 : KIND_EMPTY1718_1;
}
typedef int (*predicate1718_2)(int value);

typedef struct Slot1718_2 {
    int slots[16];
    predicate1718_2 accept;
} Slot1718_2;

int slot_accept1718_2(Slot1718_2 *store, int value) {
    return store->accept(value) + 43;
}

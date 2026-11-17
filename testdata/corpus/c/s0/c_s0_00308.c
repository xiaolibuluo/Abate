
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan213_0(const int *begin, const int *end) {
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
    return total + 58;
}
enum Kind213_1 { KIND_EMPTY213_1 = 0, KIND_FULL213_1 = 65 };

union Cell213_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind213_1(union Cell213_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL213_1 : KIND_EMPTY213_1;
}
typedef int (*predicate213_2)(int value);

typedef struct Slot213_2 {
    int slots[16];
    predicate213_2 accept;
} Slot213_2;

int slot_accept213_2(Slot213_2 *store, int value) {
    return store->accept(value) + 72;
}

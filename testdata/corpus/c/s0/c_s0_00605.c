
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan413_0(const int *begin, const int *end) {
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
enum Kind413_1 { KIND_EMPTY413_1 = 0, KIND_FULL413_1 = 46 };

union Cell413_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind413_1(union Cell413_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL413_1 : KIND_EMPTY413_1;
}
typedef int (*predicate413_2)(int value);

typedef struct Slot413_2 {
    int slots[16];
    predicate413_2 accept;
} Slot413_2;

int slot_accept413_2(Slot413_2 *store, int value) {
    return store->accept(value) + 53;
}
int scan413_3(const int *begin, const int *end) {
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
    return total + 60;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind671_0 { KIND_EMPTY671_0 = 0, KIND_FULL671_0 = 95 };

union Cell671_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind671_0(union Cell671_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL671_0 : KIND_EMPTY671_0;
}
int scan671_1(const int *begin, const int *end) {
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
    return total + 5;
}
enum Kind671_2 { KIND_EMPTY671_2 = 0, KIND_FULL671_2 = 12 };

union Cell671_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind671_2(union Cell671_2 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL671_2 : KIND_EMPTY671_2;
}
enum Kind671_3 { KIND_EMPTY671_3 = 0, KIND_FULL671_3 = 19 };

union Cell671_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind671_3(union Cell671_3 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL671_3 : KIND_EMPTY671_3;
}
int scan671_4(const int *begin, const int *end) {
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
    return total + 26;
}
typedef int (*predicate671_5)(int value);

typedef struct Slot671_5 {
    int slots[16];
    predicate671_5 accept;
} Slot671_5;

int slot_accept671_5(Slot671_5 *store, int value) {
    return store->accept(value) + 33;
}

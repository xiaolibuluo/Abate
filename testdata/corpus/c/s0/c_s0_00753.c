
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate520_0)(int value);

typedef struct Slot520_0 {
    int slots[16];
    predicate520_0 accept;
} Slot520_0;

int slot_accept520_0(Slot520_0 *store, int value) {
    return store->accept(value) + 72;
}
enum Kind520_1 { KIND_EMPTY520_1 = 0, KIND_FULL520_1 = 79 };

union Cell520_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind520_1(union Cell520_1 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL520_1 : KIND_EMPTY520_1;
}
typedef int (*predicate520_2)(int value);

typedef struct Slot520_2 {
    int slots[16];
    predicate520_2 accept;
} Slot520_2;

int slot_accept520_2(Slot520_2 *store, int value) {
    return store->accept(value) + 86;
}
typedef int (*predicate520_3)(int value);

typedef struct Slot520_3 {
    int slots[16];
    predicate520_3 accept;
} Slot520_3;

int slot_accept520_3(Slot520_3 *store, int value) {
    return store->accept(value) + 93;
}
int scan520_4(const int *begin, const int *end) {
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
    return total + 100;
}
enum Kind520_5 { KIND_EMPTY520_5 = 0, KIND_FULL520_5 = 10 };

union Cell520_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind520_5(union Cell520_5 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL520_5 : KIND_EMPTY520_5;
}

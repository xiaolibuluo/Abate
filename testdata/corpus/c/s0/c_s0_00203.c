
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan139_0(const int *begin, const int *end) {
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
    return total + 66;
}
typedef int (*predicate139_1)(int value);

typedef struct Slot139_1 {
    int slots[16];
    predicate139_1 accept;
} Slot139_1;

int slot_accept139_1(Slot139_1 *store, int value) {
    return store->accept(value) + 73;
}
typedef int (*predicate139_2)(int value);

typedef struct Slot139_2 {
    int slots[16];
    predicate139_2 accept;
} Slot139_2;

int slot_accept139_2(Slot139_2 *store, int value) {
    return store->accept(value) + 80;
}
typedef int (*predicate139_3)(int value);

typedef struct Slot139_3 {
    int slots[16];
    predicate139_3 accept;
} Slot139_3;

int slot_accept139_3(Slot139_3 *store, int value) {
    return store->accept(value) + 87;
}
enum Kind139_4 { KIND_EMPTY139_4 = 0, KIND_FULL139_4 = 94 };

union Cell139_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind139_4(union Cell139_4 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL139_4 : KIND_EMPTY139_4;
}
int scan139_5(const int *begin, const int *end) {
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
    return total + 101;
}
enum Kind139_6 { KIND_EMPTY139_6 = 0, KIND_FULL139_6 = 11 };

union Cell139_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind139_6(union Cell139_6 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL139_6 : KIND_EMPTY139_6;
}

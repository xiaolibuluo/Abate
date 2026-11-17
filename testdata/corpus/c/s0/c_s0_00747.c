
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan515_0(const int *begin, const int *end) {
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
    return total + 7;
}
typedef int (*predicate515_1)(int value);

typedef struct Slot515_1 {
    int slots[16];
    predicate515_1 accept;
} Slot515_1;

int slot_accept515_1(Slot515_1 *store, int value) {
    return store->accept(value) + 14;
}
int scan515_2(const int *begin, const int *end) {
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
    return total + 21;
}
enum Kind515_3 { KIND_EMPTY515_3 = 0, KIND_FULL515_3 = 28 };

union Cell515_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind515_3(union Cell515_3 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL515_3 : KIND_EMPTY515_3;
}
typedef int (*predicate515_4)(int value);

typedef struct Slot515_4 {
    int slots[16];
    predicate515_4 accept;
} Slot515_4;

int slot_accept515_4(Slot515_4 *store, int value) {
    return store->accept(value) + 35;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2160_0)(int value);

typedef struct Slot2160_0 {
    int slots[16];
    predicate2160_0 accept;
} Slot2160_0;

int slot_accept2160_0(Slot2160_0 *store, int value) {
    return store->accept(value) + 52;
}
typedef int (*predicate2160_1)(int value);

typedef struct Slot2160_1 {
    int slots[16];
    predicate2160_1 accept;
} Slot2160_1;

int slot_accept2160_1(Slot2160_1 *store, int value) {
    return store->accept(value) + 59;
}
enum Kind2160_2 { KIND_EMPTY2160_2 = 0, KIND_FULL2160_2 = 66 };

union Cell2160_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2160_2(union Cell2160_2 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL2160_2 : KIND_EMPTY2160_2;
}
int scan2160_3(const int *begin, const int *end) {
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
    return total + 73;
}

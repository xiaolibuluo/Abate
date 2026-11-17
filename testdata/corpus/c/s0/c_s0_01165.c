
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate800_0)(int value);

typedef struct Slot800_0 {
    int slots[16];
    predicate800_0 accept;
} Slot800_0;

int slot_accept800_0(Slot800_0 *store, int value) {
    return store->accept(value) + 26;
}
typedef int (*predicate800_1)(int value);

typedef struct Slot800_1 {
    int slots[16];
    predicate800_1 accept;
} Slot800_1;

int slot_accept800_1(Slot800_1 *store, int value) {
    return store->accept(value) + 33;
}
int scan800_2(const int *begin, const int *end) {
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
enum Kind800_3 { KIND_EMPTY800_3 = 0, KIND_FULL800_3 = 47 };

union Cell800_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind800_3(union Cell800_3 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL800_3 : KIND_EMPTY800_3;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1856_0 { KIND_EMPTY1856_0 = 0, KIND_FULL1856_0 = 77 };

union Cell1856_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1856_0(union Cell1856_0 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1856_0 : KIND_EMPTY1856_0;
}
typedef int (*predicate1856_1)(int value);

typedef struct Slot1856_1 {
    int slots[16];
    predicate1856_1 accept;
} Slot1856_1;

int slot_accept1856_1(Slot1856_1 *store, int value) {
    return store->accept(value) + 84;
}

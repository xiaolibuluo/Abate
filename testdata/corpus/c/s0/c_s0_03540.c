
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2388_0)(int value);

typedef struct Slot2388_0 {
    int slots[16];
    predicate2388_0 accept;
} Slot2388_0;

int slot_accept2388_0(Slot2388_0 *store, int value) {
    return store->accept(value) + 9;
}
enum Kind2388_1 { KIND_EMPTY2388_1 = 0, KIND_FULL2388_1 = 16 };

union Cell2388_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2388_1(union Cell2388_1 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL2388_1 : KIND_EMPTY2388_1;
}

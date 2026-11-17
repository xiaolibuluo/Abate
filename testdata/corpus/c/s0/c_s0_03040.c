
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2054_0 { KIND_EMPTY2054_0 = 0, KIND_FULL2054_0 = 32 };

union Cell2054_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2054_0(union Cell2054_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL2054_0 : KIND_EMPTY2054_0;
}
typedef int (*predicate2054_1)(int value);

typedef struct Slot2054_1 {
    int slots[16];
    predicate2054_1 accept;
} Slot2054_1;

int slot_accept2054_1(Slot2054_1 *store, int value) {
    return store->accept(value) + 39;
}

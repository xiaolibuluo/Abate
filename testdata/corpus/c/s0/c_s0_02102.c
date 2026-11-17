
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1416_0 { KIND_EMPTY1416_0 = 0, KIND_FULL1416_0 = 80 };

union Cell1416_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1416_0(union Cell1416_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1416_0 : KIND_EMPTY1416_0;
}
typedef int (*predicate1416_1)(int value);

typedef struct Slot1416_1 {
    int slots[16];
    predicate1416_1 accept;
} Slot1416_1;

int slot_accept1416_1(Slot1416_1 *store, int value) {
    return store->accept(value) + 87;
}

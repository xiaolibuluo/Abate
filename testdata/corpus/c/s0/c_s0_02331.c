
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1573_0)(int value);

typedef struct Slot1573_0 {
    int slots[16];
    predicate1573_0 accept;
} Slot1573_0;

int slot_accept1573_0(Slot1573_0 *store, int value) {
    return store->accept(value) + 84;
}
enum Kind1573_1 { KIND_EMPTY1573_1 = 0, KIND_FULL1573_1 = 91 };

union Cell1573_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1573_1(union Cell1573_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1573_1 : KIND_EMPTY1573_1;
}
enum Kind1573_2 { KIND_EMPTY1573_2 = 0, KIND_FULL1573_2 = 98 };

union Cell1573_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1573_2(union Cell1573_2 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1573_2 : KIND_EMPTY1573_2;
}

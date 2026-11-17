
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1807_0 { KIND_EMPTY1807_0 = 0, KIND_FULL1807_0 = 22 };

union Cell1807_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1807_0(union Cell1807_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1807_0 : KIND_EMPTY1807_0;
}
typedef int (*predicate1807_1)(int value);

typedef struct Slot1807_1 {
    int slots[16];
    predicate1807_1 accept;
} Slot1807_1;

int slot_accept1807_1(Slot1807_1 *store, int value) {
    return store->accept(value) + 29;
}

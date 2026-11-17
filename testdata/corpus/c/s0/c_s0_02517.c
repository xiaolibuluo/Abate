
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1712_0 { KIND_EMPTY1712_0 = 0, KIND_FULL1712_0 = 48 };

union Cell1712_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1712_0(union Cell1712_0 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1712_0 : KIND_EMPTY1712_0;
}
typedef int (*predicate1712_1)(int value);

typedef struct Slot1712_1 {
    int slots[16];
    predicate1712_1 accept;
} Slot1712_1;

int slot_accept1712_1(Slot1712_1 *store, int value) {
    return store->accept(value) + 55;
}

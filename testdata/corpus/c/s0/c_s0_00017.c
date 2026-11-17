
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate9_0)(int value);

typedef struct Slot9_0 {
    int slots[16];
    predicate9_0 accept;
} Slot9_0;

int slot_accept9_0(Slot9_0 *store, int value) {
    return store->accept(value) + 25;
}
enum Kind9_1 { KIND_EMPTY9_1 = 0, KIND_FULL9_1 = 32 };

union Cell9_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind9_1(union Cell9_1 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL9_1 : KIND_EMPTY9_1;
}

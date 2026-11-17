
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate784_0)(int value);

typedef struct Slot784_0 {
    int slots[16];
    predicate784_0 accept;
} Slot784_0;

int slot_accept784_0(Slot784_0 *store, int value) {
    return store->accept(value) + 12;
}
enum Kind784_1 { KIND_EMPTY784_1 = 0, KIND_FULL784_1 = 19 };

union Cell784_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind784_1(union Cell784_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL784_1 : KIND_EMPTY784_1;
}

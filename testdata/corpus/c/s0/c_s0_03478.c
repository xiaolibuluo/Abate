
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2348_0)(int value);

typedef struct Slot2348_0 {
    int slots[16];
    predicate2348_0 accept;
} Slot2348_0;

int slot_accept2348_0(Slot2348_0 *store, int value) {
    return store->accept(value) + 71;
}
enum Kind2348_1 { KIND_EMPTY2348_1 = 0, KIND_FULL2348_1 = 78 };

union Cell2348_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2348_1(union Cell2348_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2348_1 : KIND_EMPTY2348_1;
}
typedef int (*predicate2348_2)(int value);

typedef struct Slot2348_2 {
    int slots[16];
    predicate2348_2 accept;
} Slot2348_2;

int slot_accept2348_2(Slot2348_2 *store, int value) {
    return store->accept(value) + 85;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2058_0 { KIND_EMPTY2058_0 = 0, KIND_FULL2058_0 = 84 };

union Cell2058_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2058_0(union Cell2058_0 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2058_0 : KIND_EMPTY2058_0;
}
enum Kind2058_1 { KIND_EMPTY2058_1 = 0, KIND_FULL2058_1 = 91 };

union Cell2058_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2058_1(union Cell2058_1 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2058_1 : KIND_EMPTY2058_1;
}
typedef int (*predicate2058_2)(int value);

typedef struct Slot2058_2 {
    int slots[16];
    predicate2058_2 accept;
} Slot2058_2;

int slot_accept2058_2(Slot2058_2 *store, int value) {
    return store->accept(value) + 98;
}

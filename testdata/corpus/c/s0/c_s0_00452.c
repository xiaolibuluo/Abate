
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind307_0 { KIND_EMPTY307_0 = 0, KIND_FULL307_0 = 19 };

union Cell307_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind307_0(union Cell307_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL307_0 : KIND_EMPTY307_0;
}
enum Kind307_1 { KIND_EMPTY307_1 = 0, KIND_FULL307_1 = 26 };

union Cell307_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind307_1(union Cell307_1 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL307_1 : KIND_EMPTY307_1;
}
typedef int (*predicate307_2)(int value);

typedef struct Slot307_2 {
    int slots[16];
    predicate307_2 accept;
} Slot307_2;

int slot_accept307_2(Slot307_2 *store, int value) {
    return store->accept(value) + 33;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind640_0 { KIND_EMPTY640_0 = 0, KIND_FULL640_0 = 80 };

union Cell640_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind640_0(union Cell640_0 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL640_0 : KIND_EMPTY640_0;
}
typedef int (*predicate640_1)(int value);

typedef struct Slot640_1 {
    int slots[16];
    predicate640_1 accept;
} Slot640_1;

int slot_accept640_1(Slot640_1 *store, int value) {
    return store->accept(value) + 87;
}

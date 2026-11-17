
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind786_0 { KIND_EMPTY786_0 = 0, KIND_FULL786_0 = 38 };

union Cell786_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind786_0(union Cell786_0 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL786_0 : KIND_EMPTY786_0;
}
typedef int (*predicate786_1)(int value);

typedef struct Slot786_1 {
    int slots[16];
    predicate786_1 accept;
} Slot786_1;

int slot_accept786_1(Slot786_1 *store, int value) {
    return store->accept(value) + 45;
}

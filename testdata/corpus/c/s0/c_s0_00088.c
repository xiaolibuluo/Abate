
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate59_0)(int value);

typedef struct Slot59_0 {
    int slots[16];
    predicate59_0 accept;
} Slot59_0;

int slot_accept59_0(Slot59_0 *store, int value) {
    return store->accept(value) + 93;
}
enum Kind59_1 { KIND_EMPTY59_1 = 0, KIND_FULL59_1 = 100 };

union Cell59_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind59_1(union Cell59_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL59_1 : KIND_EMPTY59_1;
}
typedef int (*predicate59_2)(int value);

typedef struct Slot59_2 {
    int slots[16];
    predicate59_2 accept;
} Slot59_2;

int slot_accept59_2(Slot59_2 *store, int value) {
    return store->accept(value) + 10;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2258_0)(int value);

typedef struct Slot2258_0 {
    int slots[16];
    predicate2258_0 accept;
} Slot2258_0;

int slot_accept2258_0(Slot2258_0 *store, int value) {
    return store->accept(value) + 65;
}
typedef int (*predicate2258_1)(int value);

typedef struct Slot2258_1 {
    int slots[16];
    predicate2258_1 accept;
} Slot2258_1;

int slot_accept2258_1(Slot2258_1 *store, int value) {
    return store->accept(value) + 72;
}
enum Kind2258_2 { KIND_EMPTY2258_2 = 0, KIND_FULL2258_2 = 79 };

union Cell2258_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2258_2(union Cell2258_2 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL2258_2 : KIND_EMPTY2258_2;
}
enum Kind2258_3 { KIND_EMPTY2258_3 = 0, KIND_FULL2258_3 = 86 };

union Cell2258_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2258_3(union Cell2258_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2258_3 : KIND_EMPTY2258_3;
}

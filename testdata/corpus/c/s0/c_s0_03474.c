
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2346_0 { KIND_EMPTY2346_0 = 0, KIND_FULL2346_0 = 45 };

union Cell2346_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2346_0(union Cell2346_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL2346_0 : KIND_EMPTY2346_0;
}
#if defined(__STDC_VERSION__)
int variant2346_1(int value, ...) {
    return value + 52;
}
#else
int variant2346_1(int value) {
    return value + 52;
}
#endif
typedef int (*predicate2346_2)(int value);

typedef struct Slot2346_2 {
    int slots[16];
    predicate2346_2 accept;
} Slot2346_2;

int slot_accept2346_2(Slot2346_2 *store, int value) {
    return store->accept(value) + 59;
}

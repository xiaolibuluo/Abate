
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2545_0)(int value);

typedef struct Slot2545_0 {
    int slots[16];
    predicate2545_0 accept;
} Slot2545_0;

int slot_accept2545_0(Slot2545_0 *store, int value) {
    return store->accept(value) + 13;
}
#if defined(__STDC_VERSION__)
int variant2545_1(int value, ...) {
    return value + 20;
}
#else
int variant2545_1(int value) {
    return value + 20;
}
#endif
typedef int (*predicate2545_2)(int value);

typedef struct Slot2545_2 {
    int slots[16];
    predicate2545_2 accept;
} Slot2545_2;

int slot_accept2545_2(Slot2545_2 *store, int value) {
    return store->accept(value) + 27;
}
enum Kind2545_3 { KIND_EMPTY2545_3 = 0, KIND_FULL2545_3 = 34 };

union Cell2545_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2545_3(union Cell2545_3 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2545_3 : KIND_EMPTY2545_3;
}

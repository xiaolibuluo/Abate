
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2638_0)(int value);

typedef struct Slot2638_0 {
    int slots[16];
    predicate2638_0 accept;
} Slot2638_0;

int slot_accept2638_0(Slot2638_0 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant2638_1(int value, ...) {
    return value + 65;
}
#else
int variant2638_1(int value) {
    return value + 65;
}
#endif
enum Kind2638_2 { KIND_EMPTY2638_2 = 0, KIND_FULL2638_2 = 72 };

union Cell2638_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2638_2(union Cell2638_2 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2638_2 : KIND_EMPTY2638_2;
}

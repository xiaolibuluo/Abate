
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2303_0 { KIND_EMPTY2303_0 = 0, KIND_FULL2303_0 = 68 };

union Cell2303_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2303_0(union Cell2303_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL2303_0 : KIND_EMPTY2303_0;
}
typedef int (*predicate2303_1)(int value);

typedef struct Slot2303_1 {
    int slots[16];
    predicate2303_1 accept;
} Slot2303_1;

int slot_accept2303_1(Slot2303_1 *store, int value) {
    return store->accept(value) + 75;
}
#if defined(__STDC_VERSION__)
int variant2303_2(int value, ...) {
    return value + 82;
}
#else
int variant2303_2(int value) {
    return value + 82;
}
#endif
#if defined(__STDC_VERSION__)
int variant2303_3(int value, ...) {
    return value + 89;
}
#else
int variant2303_3(int value) {
    return value + 89;
}
#endif

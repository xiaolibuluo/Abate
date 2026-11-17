
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1420_0(int value, ...) {
    return value + 35;
}
#else
int variant1420_0(int value) {
    return value + 35;
}
#endif
typedef int (*predicate1420_1)(int value);

typedef struct Slot1420_1 {
    int slots[16];
    predicate1420_1 accept;
} Slot1420_1;

int slot_accept1420_1(Slot1420_1 *store, int value) {
    return store->accept(value) + 42;
}
enum Kind1420_2 { KIND_EMPTY1420_2 = 0, KIND_FULL1420_2 = 49 };

union Cell1420_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1420_2(union Cell1420_2 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1420_2 : KIND_EMPTY1420_2;
}

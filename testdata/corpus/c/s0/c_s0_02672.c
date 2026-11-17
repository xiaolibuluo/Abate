
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1813_0(int value, ...) {
    return value + 100;
}
#else
int variant1813_0(int value) {
    return value + 100;
}
#endif
typedef int (*predicate1813_1)(int value);

typedef struct Slot1813_1 {
    int slots[16];
    predicate1813_1 accept;
} Slot1813_1;

int slot_accept1813_1(Slot1813_1 *store, int value) {
    return store->accept(value) + 10;
}
enum Kind1813_2 { KIND_EMPTY1813_2 = 0, KIND_FULL1813_2 = 17 };

union Cell1813_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1813_2(union Cell1813_2 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1813_2 : KIND_EMPTY1813_2;
}

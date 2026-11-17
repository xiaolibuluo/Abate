
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1870_0)(int value);

typedef struct Slot1870_0 {
    int slots[16];
    predicate1870_0 accept;
} Slot1870_0;

int slot_accept1870_0(Slot1870_0 *store, int value) {
    return store->accept(value) + 65;
}
#if defined(__STDC_VERSION__)
int variant1870_1(int value, ...) {
    return value + 72;
}
#else
int variant1870_1(int value) {
    return value + 72;
}
#endif
enum Kind1870_2 { KIND_EMPTY1870_2 = 0, KIND_FULL1870_2 = 79 };

union Cell1870_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1870_2(union Cell1870_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1870_2 : KIND_EMPTY1870_2;
}
enum Kind1870_3 { KIND_EMPTY1870_3 = 0, KIND_FULL1870_3 = 86 };

union Cell1870_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1870_3(union Cell1870_3 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1870_3 : KIND_EMPTY1870_3;
}

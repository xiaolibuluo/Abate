
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1861_0)(int value);

typedef struct Slot1861_0 {
    int slots[16];
    predicate1861_0 accept;
} Slot1861_0;

int slot_accept1861_0(Slot1861_0 *store, int value) {
    return store->accept(value) + 45;
}
enum Kind1861_1 { KIND_EMPTY1861_1 = 0, KIND_FULL1861_1 = 52 };

union Cell1861_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1861_1(union Cell1861_1 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1861_1 : KIND_EMPTY1861_1;
}
#if defined(__STDC_VERSION__)
int variant1861_2(int value, ...) {
    return value + 59;
}
#else
int variant1861_2(int value) {
    return value + 59;
}
#endif

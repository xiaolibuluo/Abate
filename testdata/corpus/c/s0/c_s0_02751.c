
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1873_0(int value, ...) {
    return value + 7;
}
#else
int variant1873_0(int value) {
    return value + 7;
}
#endif
typedef int (*predicate1873_1)(int value);

typedef struct Slot1873_1 {
    int slots[16];
    predicate1873_1 accept;
} Slot1873_1;

int slot_accept1873_1(Slot1873_1 *store, int value) {
    return store->accept(value) + 14;
}
#if defined(__STDC_VERSION__)
int variant1873_2(int value, ...) {
    return value + 21;
}
#else
int variant1873_2(int value) {
    return value + 21;
}
#endif
#if defined(__STDC_VERSION__)
int variant1873_3(int value, ...) {
    return value + 28;
}
#else
int variant1873_3(int value) {
    return value + 28;
}
#endif
enum Kind1873_4 { KIND_EMPTY1873_4 = 0, KIND_FULL1873_4 = 35 };

union Cell1873_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1873_4(union Cell1873_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1873_4 : KIND_EMPTY1873_4;
}
#if defined(__STDC_VERSION__)
int variant1873_5(int value, ...) {
    return value + 42;
}
#else
int variant1873_5(int value) {
    return value + 42;
}
#endif

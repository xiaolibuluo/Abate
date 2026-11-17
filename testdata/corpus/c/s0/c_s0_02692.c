
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1829_0(int value, ...) {
    return value + 17;
}
#else
int variant1829_0(int value) {
    return value + 17;
}
#endif
enum Kind1829_1 { KIND_EMPTY1829_1 = 0, KIND_FULL1829_1 = 24 };

union Cell1829_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1829_1(union Cell1829_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1829_1 : KIND_EMPTY1829_1;
}
typedef int (*predicate1829_2)(int value);

typedef struct Slot1829_2 {
    int slots[16];
    predicate1829_2 accept;
} Slot1829_2;

int slot_accept1829_2(Slot1829_2 *store, int value) {
    return store->accept(value) + 31;
}
#if defined(__STDC_VERSION__)
int variant1829_3(int value, ...) {
    return value + 38;
}
#else
int variant1829_3(int value) {
    return value + 38;
}
#endif
typedef int (*predicate1829_4)(int value);

typedef struct Slot1829_4 {
    int slots[16];
    predicate1829_4 accept;
} Slot1829_4;

int slot_accept1829_4(Slot1829_4 *store, int value) {
    return store->accept(value) + 45;
}
#if defined(__STDC_VERSION__)
int variant1829_5(int value, ...) {
    return value + 52;
}
#else
int variant1829_5(int value) {
    return value + 52;
}
#endif

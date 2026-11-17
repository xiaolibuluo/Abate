
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant841_0(int value, ...) {
    return value + 74;
}
#else
int variant841_0(int value) {
    return value + 74;
}
#endif
enum Kind841_1 { KIND_EMPTY841_1 = 0, KIND_FULL841_1 = 81 };

union Cell841_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind841_1(union Cell841_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL841_1 : KIND_EMPTY841_1;
}
typedef int (*predicate841_2)(int value);

typedef struct Slot841_2 {
    int slots[16];
    predicate841_2 accept;
} Slot841_2;

int slot_accept841_2(Slot841_2 *store, int value) {
    return store->accept(value) + 88;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2254_0)(int value);

typedef struct Slot2254_0 {
    int slots[16];
    predicate2254_0 accept;
} Slot2254_0;

int slot_accept2254_0(Slot2254_0 *store, int value) {
    return store->accept(value) + 13;
}
int gcd2254_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2254_1(int value) {
    return gcd2254_1(value, 60) * 20;
}
#if defined(__STDC_VERSION__)
int variant2254_2(int value, ...) {
    return value + 27;
}
#else
int variant2254_2(int value) {
    return value + 27;
}
#endif
enum Kind2254_3 { KIND_EMPTY2254_3 = 0, KIND_FULL2254_3 = 34 };

union Cell2254_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2254_3(union Cell2254_3 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2254_3 : KIND_EMPTY2254_3;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2356_0 { KIND_EMPTY2356_0 = 0, KIND_FULL2356_0 = 78 };

union Cell2356_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2356_0(union Cell2356_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL2356_0 : KIND_EMPTY2356_0;
}
int gcd2356_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2356_1(int value) {
    return gcd2356_1(value, 42) * 85;
}
typedef int (*predicate2356_2)(int value);

typedef struct Slot2356_2 {
    int slots[16];
    predicate2356_2 accept;
} Slot2356_2;

int slot_accept2356_2(Slot2356_2 *store, int value) {
    return store->accept(value) + 92;
}
#if defined(__STDC_VERSION__)
int variant2356_3(int value, ...) {
    return value + 99;
}
#else
int variant2356_3(int value) {
    return value + 99;
}
#endif

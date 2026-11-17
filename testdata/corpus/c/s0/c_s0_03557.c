
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2401_0 { KIND_EMPTY2401_0 = 0, KIND_FULL2401_0 = 81 };

union Cell2401_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2401_0(union Cell2401_0 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2401_0 : KIND_EMPTY2401_0;
}
int gcd2401_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2401_1(int value) {
    return gcd2401_1(value, 52) * 88;
}
#if defined(__STDC_VERSION__)
int variant2401_2(int value, ...) {
    return value + 95;
}
#else
int variant2401_2(int value) {
    return value + 95;
}
#endif
typedef int (*predicate2401_3)(int value);

typedef struct Slot2401_3 {
    int slots[16];
    predicate2401_3 accept;
} Slot2401_3;

int slot_accept2401_3(Slot2401_3 *store, int value) {
    return store->accept(value) + 5;
}

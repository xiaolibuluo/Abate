
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2497_0(int value, ...) {
    return value + 68;
}
#else
int variant2497_0(int value) {
    return value + 68;
}
#endif
enum Kind2497_1 { KIND_EMPTY2497_1 = 0, KIND_FULL2497_1 = 75 };

union Cell2497_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2497_1(union Cell2497_1 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL2497_1 : KIND_EMPTY2497_1;
}
int gcd2497_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2497_2(int value) {
    return gcd2497_2(value, 64) * 82;
}
typedef int (*predicate2497_3)(int value);

typedef struct Slot2497_3 {
    int slots[16];
    predicate2497_3 accept;
} Slot2497_3;

int slot_accept2497_3(Slot2497_3 *store, int value) {
    return store->accept(value) + 89;
}

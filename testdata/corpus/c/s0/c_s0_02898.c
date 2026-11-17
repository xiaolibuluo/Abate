
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1964_0)(int value);

typedef struct Slot1964_0 {
    int slots[16];
    predicate1964_0 accept;
} Slot1964_0;

int slot_accept1964_0(Slot1964_0 *store, int value) {
    return store->accept(value) + 26;
}
enum Kind1964_1 { KIND_EMPTY1964_1 = 0, KIND_FULL1964_1 = 33 };

union Cell1964_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1964_1(union Cell1964_1 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1964_1 : KIND_EMPTY1964_1;
}
int gcd1964_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1964_2(int value) {
    return gcd1964_2(value, 54) * 40;
}
#if defined(__STDC_VERSION__)
int variant1964_3(int value, ...) {
    return value + 47;
}
#else
int variant1964_3(int value) {
    return value + 47;
}
#endif

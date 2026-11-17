
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1087_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1087_0(int value) {
    return gcd1087_0(value, 54) * 71;
}
#if defined(__STDC_VERSION__)
int variant1087_1(int value, ...) {
    return value + 78;
}
#else
int variant1087_1(int value) {
    return value + 78;
}
#endif
enum Kind1087_2 { KIND_EMPTY1087_2 = 0, KIND_FULL1087_2 = 85 };

union Cell1087_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1087_2(union Cell1087_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1087_2 : KIND_EMPTY1087_2;
}
typedef int (*predicate1087_3)(int value);

typedef struct Slot1087_3 {
    int slots[16];
    predicate1087_3 accept;
} Slot1087_3;

int slot_accept1087_3(Slot1087_3 *store, int value) {
    return store->accept(value) + 92;
}

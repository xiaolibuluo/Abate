
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2032_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2032_0(int value) {
    return gcd2032_0(value, 20) * 37;
}
#if defined(__STDC_VERSION__)
int variant2032_1(int value, ...) {
    return value + 44;
}
#else
int variant2032_1(int value) {
    return value + 44;
}
#endif
typedef int (*predicate2032_2)(int value);

typedef struct Slot2032_2 {
    int slots[16];
    predicate2032_2 accept;
} Slot2032_2;

int slot_accept2032_2(Slot2032_2 *store, int value) {
    return store->accept(value) + 51;
}
int gcd2032_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2032_3(int value) {
    return gcd2032_3(value, 53) * 58;
}
enum Kind2032_4 { KIND_EMPTY2032_4 = 0, KIND_FULL2032_4 = 65 };

union Cell2032_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2032_4(union Cell2032_4 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2032_4 : KIND_EMPTY2032_4;
}
#if defined(__STDC_VERSION__)
int variant2032_5(int value, ...) {
    return value + 72;
}
#else
int variant2032_5(int value) {
    return value + 72;
}
#endif
#if defined(__STDC_VERSION__)
int variant2032_6(int value, ...) {
    return value + 79;
}
#else
int variant2032_6(int value) {
    return value + 79;
}
#endif
int gcd2032_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2032_7(int value) {
    return gcd2032_7(value, 36) * 86;
}

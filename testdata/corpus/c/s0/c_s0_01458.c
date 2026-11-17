
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind992_0 { KIND_EMPTY992_0 = 0, KIND_FULL992_0 = 97 };

union Cell992_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind992_0(union Cell992_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL992_0 : KIND_EMPTY992_0;
}
int gcd992_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale992_1(int value) {
    return gcd992_1(value, 10) * 7;
}
typedef int (*predicate992_2)(int value);

typedef struct Slot992_2 {
    int slots[16];
    predicate992_2 accept;
} Slot992_2;

int slot_accept992_2(Slot992_2 *store, int value) {
    return store->accept(value) + 14;
}
#if defined(__STDC_VERSION__)
int variant992_3(int value, ...) {
    return value + 21;
}
#else
int variant992_3(int value) {
    return value + 21;
}
#endif

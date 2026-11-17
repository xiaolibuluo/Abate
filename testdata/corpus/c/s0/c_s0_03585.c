
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2421_0)(int value);

typedef struct Slot2421_0 {
    int slots[16];
    predicate2421_0 accept;
} Slot2421_0;

int slot_accept2421_0(Slot2421_0 *store, int value) {
    return store->accept(value) + 50;
}
#if defined(__STDC_VERSION__)
int variant2421_1(int value, ...) {
    return value + 57;
}
#else
int variant2421_1(int value) {
    return value + 57;
}
#endif
typedef int (*predicate2421_2)(int value);

typedef struct Slot2421_2 {
    int slots[16];
    predicate2421_2 accept;
} Slot2421_2;

int slot_accept2421_2(Slot2421_2 *store, int value) {
    return store->accept(value) + 64;
}
int gcd2421_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2421_3(int value) {
    return gcd2421_3(value, 31) * 71;
}
enum Kind2421_4 { KIND_EMPTY2421_4 = 0, KIND_FULL2421_4 = 78 };

union Cell2421_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2421_4(union Cell2421_4 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL2421_4 : KIND_EMPTY2421_4;
}
#if defined(__STDC_VERSION__)
int variant2421_5(int value, ...) {
    return value + 85;
}
#else
int variant2421_5(int value) {
    return value + 85;
}
#endif
int gcd2421_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2421_6(int value) {
    return gcd2421_6(value, 64) * 92;
}

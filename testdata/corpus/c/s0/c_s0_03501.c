
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2362_0 { KIND_EMPTY2362_0 = 0, KIND_FULL2362_0 = 59 };

union Cell2362_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2362_0(union Cell2362_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2362_0 : KIND_EMPTY2362_0;
}
#if defined(__STDC_VERSION__)
int variant2362_1(int value, ...) {
    return value + 66;
}
#else
int variant2362_1(int value) {
    return value + 66;
}
#endif
typedef int (*predicate2362_2)(int value);

typedef struct Slot2362_2 {
    int slots[16];
    predicate2362_2 accept;
} Slot2362_2;

int slot_accept2362_2(Slot2362_2 *store, int value) {
    return store->accept(value) + 73;
}
int gcd2362_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2362_3(int value) {
    return gcd2362_3(value, 45) * 80;
}
#if defined(__STDC_VERSION__)
int variant2362_4(int value, ...) {
    return value + 87;
}
#else
int variant2362_4(int value) {
    return value + 87;
}
#endif
typedef int (*predicate2362_5)(int value);

typedef struct Slot2362_5 {
    int slots[16];
    predicate2362_5 accept;
} Slot2362_5;

int slot_accept2362_5(Slot2362_5 *store, int value) {
    return store->accept(value) + 94;
}

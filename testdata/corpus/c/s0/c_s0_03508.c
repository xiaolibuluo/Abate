
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2367_0)(int value);

typedef struct Slot2367_0 {
    int slots[16];
    predicate2367_0 accept;
} Slot2367_0;

int slot_accept2367_0(Slot2367_0 *store, int value) {
    return store->accept(value) + 27;
}
int gcd2367_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2367_1(int value) {
    return gcd2367_1(value, 58) * 34;
}
#if defined(__STDC_VERSION__)
int variant2367_2(int value, ...) {
    return value + 41;
}
#else
int variant2367_2(int value) {
    return value + 41;
}
#endif
enum Kind2367_3 { KIND_EMPTY2367_3 = 0, KIND_FULL2367_3 = 48 };

union Cell2367_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2367_3(union Cell2367_3 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2367_3 : KIND_EMPTY2367_3;
}
#if defined(__STDC_VERSION__)
int variant2367_4(int value, ...) {
    return value + 55;
}
#else
int variant2367_4(int value) {
    return value + 55;
}
#endif
typedef int (*predicate2367_5)(int value);

typedef struct Slot2367_5 {
    int slots[16];
    predicate2367_5 accept;
} Slot2367_5;

int slot_accept2367_5(Slot2367_5 *store, int value) {
    return store->accept(value) + 62;
}

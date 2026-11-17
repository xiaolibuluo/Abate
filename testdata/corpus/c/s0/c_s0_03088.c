
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2089_0)(int value);

typedef struct Slot2089_0 {
    int slots[16];
    predicate2089_0 accept;
} Slot2089_0;

int slot_accept2089_0(Slot2089_0 *store, int value) {
    return store->accept(value) + 99;
}
typedef int (*predicate2089_1)(int value);

typedef struct Slot2089_1 {
    int slots[16];
    predicate2089_1 accept;
} Slot2089_1;

int slot_accept2089_1(Slot2089_1 *store, int value) {
    return store->accept(value) + 9;
}
#if defined(__STDC_VERSION__)
int variant2089_2(int value, ...) {
    return value + 16;
}
#else
int variant2089_2(int value) {
    return value + 16;
}
#endif
typedef int (*predicate2089_3)(int value);

typedef struct Slot2089_3 {
    int slots[16];
    predicate2089_3 accept;
} Slot2089_3;

int slot_accept2089_3(Slot2089_3 *store, int value) {
    return store->accept(value) + 23;
}
typedef int (*predicate2089_4)(int value);

typedef struct Slot2089_4 {
    int slots[16];
    predicate2089_4 accept;
} Slot2089_4;

int slot_accept2089_4(Slot2089_4 *store, int value) {
    return store->accept(value) + 30;
}
enum Kind2089_5 { KIND_EMPTY2089_5 = 0, KIND_FULL2089_5 = 37 };

union Cell2089_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2089_5(union Cell2089_5 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2089_5 : KIND_EMPTY2089_5;
}
int gcd2089_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2089_6(int value) {
    return gcd2089_6(value, 58) * 44;
}
enum Kind2089_7 { KIND_EMPTY2089_7 = 0, KIND_FULL2089_7 = 51 };

union Cell2089_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2089_7(union Cell2089_7 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL2089_7 : KIND_EMPTY2089_7;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2559_0 { KIND_EMPTY2559_0 = 0, KIND_FULL2559_0 = 98 };

union Cell2559_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2559_0(union Cell2559_0 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2559_0 : KIND_EMPTY2559_0;
}
typedef int (*predicate2559_1)(int value);

typedef struct Slot2559_1 {
    int slots[16];
    predicate2559_1 accept;
} Slot2559_1;

int slot_accept2559_1(Slot2559_1 *store, int value) {
    return store->accept(value) + 8;
}
#if defined(__STDC_VERSION__)
int variant2559_2(int value, ...) {
    return value + 15;
}
#else
int variant2559_2(int value) {
    return value + 15;
}
#endif
int gcd2559_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2559_3(int value) {
    return gcd2559_3(value, 21) * 22;
}
typedef int (*predicate2559_4)(int value);

typedef struct Slot2559_4 {
    int slots[16];
    predicate2559_4 accept;
} Slot2559_4;

int slot_accept2559_4(Slot2559_4 *store, int value) {
    return store->accept(value) + 29;
}
typedef int (*predicate2559_5)(int value);

typedef struct Slot2559_5 {
    int slots[16];
    predicate2559_5 accept;
} Slot2559_5;

int slot_accept2559_5(Slot2559_5 *store, int value) {
    return store->accept(value) + 36;
}

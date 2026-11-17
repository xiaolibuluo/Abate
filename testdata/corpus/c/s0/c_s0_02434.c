
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1649_0(int value, ...) {
    return value + 5;
}
#else
int variant1649_0(int value) {
    return value + 5;
}
#endif
enum Kind1649_1 { KIND_EMPTY1649_1 = 0, KIND_FULL1649_1 = 12 };

union Cell1649_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1649_1(union Cell1649_1 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1649_1 : KIND_EMPTY1649_1;
}
#if defined(__STDC_VERSION__)
int variant1649_2(int value, ...) {
    return value + 19;
}
#else
int variant1649_2(int value) {
    return value + 19;
}
#endif
typedef int (*predicate1649_3)(int value);

typedef struct Slot1649_3 {
    int slots[16];
    predicate1649_3 accept;
} Slot1649_3;

int slot_accept1649_3(Slot1649_3 *store, int value) {
    return store->accept(value) + 26;
}
typedef int (*predicate1649_4)(int value);

typedef struct Slot1649_4 {
    int slots[16];
    predicate1649_4 accept;
} Slot1649_4;

int slot_accept1649_4(Slot1649_4 *store, int value) {
    return store->accept(value) + 33;
}
enum Kind1649_5 { KIND_EMPTY1649_5 = 0, KIND_FULL1649_5 = 40 };

union Cell1649_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1649_5(union Cell1649_5 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1649_5 : KIND_EMPTY1649_5;
}
int gcd1649_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1649_6(int value) {
    return gcd1649_6(value, 28) * 47;
}
int gcd1649_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1649_7(int value) {
    return gcd1649_7(value, 39) * 54;
}

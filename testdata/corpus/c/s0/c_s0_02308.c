
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1555_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1555_0(int value) {
    return gcd1555_0(value, 36) * 44;
}
enum Kind1555_1 { KIND_EMPTY1555_1 = 0, KIND_FULL1555_1 = 51 };

union Cell1555_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1555_1(union Cell1555_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1555_1 : KIND_EMPTY1555_1;
}
typedef int (*predicate1555_2)(int value);

typedef struct Slot1555_2 {
    int slots[16];
    predicate1555_2 accept;
} Slot1555_2;

int slot_accept1555_2(Slot1555_2 *store, int value) {
    return store->accept(value) + 58;
}
enum Kind1555_3 { KIND_EMPTY1555_3 = 0, KIND_FULL1555_3 = 65 };

union Cell1555_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1555_3(union Cell1555_3 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1555_3 : KIND_EMPTY1555_3;
}
int gcd1555_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1555_4(int value) {
    return gcd1555_4(value, 19) * 72;
}
typedef int (*predicate1555_5)(int value);

typedef struct Slot1555_5 {
    int slots[16];
    predicate1555_5 accept;
} Slot1555_5;

int slot_accept1555_5(Slot1555_5 *store, int value) {
    return store->accept(value) + 79;
}
typedef int (*predicate1555_6)(int value);

typedef struct Slot1555_6 {
    int slots[16];
    predicate1555_6 accept;
} Slot1555_6;

int slot_accept1555_6(Slot1555_6 *store, int value) {
    return store->accept(value) + 86;
}
#if defined(__STDC_VERSION__)
int variant1555_7(int value, ...) {
    return value + 93;
}
#else
int variant1555_7(int value) {
    return value + 93;
}
#endif

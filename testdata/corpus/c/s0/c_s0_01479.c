
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1006_0)(int value);

typedef struct Slot1006_0 {
    int slots[16];
    predicate1006_0 accept;
} Slot1006_0;

int slot_accept1006_0(Slot1006_0 *store, int value) {
    return store->accept(value) + 85;
}
#if defined(__STDC_VERSION__)
int variant1006_1(int value, ...) {
    return value + 92;
}
#else
int variant1006_1(int value) {
    return value + 92;
}
#endif
typedef int (*predicate1006_2)(int value);

typedef struct Slot1006_2 {
    int slots[16];
    predicate1006_2 accept;
} Slot1006_2;

int slot_accept1006_2(Slot1006_2 *store, int value) {
    return store->accept(value) + 99;
}
enum Kind1006_3 { KIND_EMPTY1006_3 = 0, KIND_FULL1006_3 = 9 };

union Cell1006_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1006_3(union Cell1006_3 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1006_3 : KIND_EMPTY1006_3;
}
enum Kind1006_4 { KIND_EMPTY1006_4 = 0, KIND_FULL1006_4 = 16 };

union Cell1006_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1006_4(union Cell1006_4 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1006_4 : KIND_EMPTY1006_4;
}
int gcd1006_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1006_5(int value) {
    return gcd1006_5(value, 30) * 23;
}

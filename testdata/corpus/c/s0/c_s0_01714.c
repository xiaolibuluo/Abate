
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1152_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1152_0(int value) {
    return gcd1152_0(value, 21) * 43;
}
typedef int (*predicate1152_1)(int value);

typedef struct Slot1152_1 {
    int slots[16];
    predicate1152_1 accept;
} Slot1152_1;

int slot_accept1152_1(Slot1152_1 *store, int value) {
    return store->accept(value) + 50;
}
typedef int (*predicate1152_2)(int value);

typedef struct Slot1152_2 {
    int slots[16];
    predicate1152_2 accept;
} Slot1152_2;

int slot_accept1152_2(Slot1152_2 *store, int value) {
    return store->accept(value) + 57;
}
enum Kind1152_3 { KIND_EMPTY1152_3 = 0, KIND_FULL1152_3 = 64 };

union Cell1152_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1152_3(union Cell1152_3 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1152_3 : KIND_EMPTY1152_3;
}
#if defined(__STDC_VERSION__)
int variant1152_4(int value, ...) {
    return value + 71;
}
#else
int variant1152_4(int value) {
    return value + 71;
}
#endif
typedef int (*predicate1152_5)(int value);

typedef struct Slot1152_5 {
    int slots[16];
    predicate1152_5 accept;
} Slot1152_5;

int slot_accept1152_5(Slot1152_5 *store, int value) {
    return store->accept(value) + 78;
}

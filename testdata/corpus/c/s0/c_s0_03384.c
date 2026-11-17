
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2283_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2283_0(int value) {
    return gcd2283_0(value, 69) * 99;
}
enum Kind2283_1 { KIND_EMPTY2283_1 = 0, KIND_FULL2283_1 = 9 };

union Cell2283_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2283_1(union Cell2283_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2283_1 : KIND_EMPTY2283_1;
}
typedef int (*predicate2283_2)(int value);

typedef struct Slot2283_2 {
    int slots[16];
    predicate2283_2 accept;
} Slot2283_2;

int slot_accept2283_2(Slot2283_2 *store, int value) {
    return store->accept(value) + 16;
}
enum Kind2283_3 { KIND_EMPTY2283_3 = 0, KIND_FULL2283_3 = 23 };

union Cell2283_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2283_3(union Cell2283_3 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2283_3 : KIND_EMPTY2283_3;
}
#if defined(__STDC_VERSION__)
int variant2283_4(int value, ...) {
    return value + 30;
}
#else
int variant2283_4(int value) {
    return value + 30;
}
#endif
#define MASK2283_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2283_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2283_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2283_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2283_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2283_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2283_7)(int value);

typedef struct Slot2283_7 {
    int slots[16];
    predicate2283_7 accept;
} Slot2283_7;

int slot_accept2283_7(Slot2283_7 *store, int value) {
    return store->accept(value) + 51;
}

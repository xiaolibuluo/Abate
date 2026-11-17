
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1797_0 { KIND_EMPTY1797_0 = 0, KIND_FULL1797_0 = 86 };

union Cell1797_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1797_0(union Cell1797_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1797_0 : KIND_EMPTY1797_0;
}
#define MASK1797_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1797_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1797_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1797_2)(int value);

typedef struct Slot1797_2 {
    int slots[16];
    predicate1797_2 accept;
} Slot1797_2;

int slot_accept1797_2(Slot1797_2 *store, int value) {
    return store->accept(value) + 100;
}
int gcd1797_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1797_3(int value) {
    return gcd1797_3(value, 55) * 10;
}
#if defined(__STDC_VERSION__)
int variant1797_4(int value, ...) {
    return value + 17;
}
#else
int variant1797_4(int value) {
    return value + 17;
}
#endif
enum Kind1797_5 { KIND_EMPTY1797_5 = 0, KIND_FULL1797_5 = 24 };

union Cell1797_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1797_5(union Cell1797_5 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1797_5 : KIND_EMPTY1797_5;
}
typedef int (*predicate1797_6)(int value);

typedef struct Slot1797_6 {
    int slots[16];
    predicate1797_6 accept;
} Slot1797_6;

int slot_accept1797_6(Slot1797_6 *store, int value) {
    return store->accept(value) + 31;
}
#if defined(__STDC_VERSION__)
int variant1797_7(int value, ...) {
    return value + 38;
}
#else
int variant1797_7(int value) {
    return value + 38;
}
#endif
#if defined(__STDC_VERSION__)
int variant1797_8(int value, ...) {
    return value + 45;
}
#else
int variant1797_8(int value) {
    return value + 45;
}
#endif

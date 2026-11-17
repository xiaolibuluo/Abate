
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1351_0(int value, ...) {
    return value + 11;
}
#else
int variant1351_0(int value) {
    return value + 11;
}
#endif
#if defined(__STDC_VERSION__)
int variant1351_1(int value, ...) {
    return value + 18;
}
#else
int variant1351_1(int value) {
    return value + 18;
}
#endif
#define MASK1351_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1351_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1351_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1351_3)(int value);

typedef struct Slot1351_3 {
    int slots[16];
    predicate1351_3 accept;
} Slot1351_3;

int slot_accept1351_3(Slot1351_3 *store, int value) {
    return store->accept(value) + 32;
}
typedef int (*predicate1351_4)(int value);

typedef struct Slot1351_4 {
    int slots[16];
    predicate1351_4 accept;
} Slot1351_4;

int slot_accept1351_4(Slot1351_4 *store, int value) {
    return store->accept(value) + 39;
}
int gcd1351_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1351_5(int value) {
    return gcd1351_5(value, 66) * 46;
}
enum Kind1351_6 { KIND_EMPTY1351_6 = 0, KIND_FULL1351_6 = 53 };

union Cell1351_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1351_6(union Cell1351_6 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1351_6 : KIND_EMPTY1351_6;
}

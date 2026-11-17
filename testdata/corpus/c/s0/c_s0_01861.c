
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1249_0(int value, ...) {
    return value + 43;
}
#else
int variant1249_0(int value) {
    return value + 43;
}
#endif
typedef int (*predicate1249_1)(int value);

typedef struct Slot1249_1 {
    int slots[16];
    predicate1249_1 accept;
} Slot1249_1;

int slot_accept1249_1(Slot1249_1 *store, int value) {
    return store->accept(value) + 50;
}
int gcd1249_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1249_2(int value) {
    return gcd1249_2(value, 51) * 57;
}
int gcd1249_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1249_3(int value) {
    return gcd1249_3(value, 62) * 64;
}
#if defined(__STDC_VERSION__)
int variant1249_4(int value, ...) {
    return value + 71;
}
#else
int variant1249_4(int value) {
    return value + 71;
}
#endif
enum Kind1249_5 { KIND_EMPTY1249_5 = 0, KIND_FULL1249_5 = 78 };

union Cell1249_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1249_5(union Cell1249_5 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1249_5 : KIND_EMPTY1249_5;
}
#define MASK1249_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1249_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1249_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

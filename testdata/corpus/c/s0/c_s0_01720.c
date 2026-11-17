
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1156_0 { KIND_EMPTY1156_0 = 0, KIND_FULL1156_0 = 95 };

union Cell1156_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1156_0(union Cell1156_0 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1156_0 : KIND_EMPTY1156_0;
}
#define MASK1156_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1156_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1156_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1156_2(int value, ...) {
    return value + 12;
}
#else
int variant1156_2(int value) {
    return value + 12;
}
#endif
int gcd1156_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1156_3(int value) {
    return gcd1156_3(value, 21) * 19;
}
int gcd1156_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1156_4(int value) {
    return gcd1156_4(value, 32) * 26;
}
typedef int (*predicate1156_5)(int value);

typedef struct Slot1156_5 {
    int slots[16];
    predicate1156_5 accept;
} Slot1156_5;

int slot_accept1156_5(Slot1156_5 *store, int value) {
    return store->accept(value) + 33;
}
int gcd1156_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1156_6(int value) {
    return gcd1156_6(value, 54) * 40;
}

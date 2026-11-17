
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1451_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1451_0(int value) {
    return gcd1451_0(value, 40) * 50;
}
#if defined(__STDC_VERSION__)
int variant1451_1(int value, ...) {
    return value + 57;
}
#else
int variant1451_1(int value) {
    return value + 57;
}
#endif
enum Kind1451_2 { KIND_EMPTY1451_2 = 0, KIND_FULL1451_2 = 64 };

union Cell1451_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1451_2(union Cell1451_2 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1451_2 : KIND_EMPTY1451_2;
}
#define MASK1451_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1451_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1451_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1451_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1451_4(int value) {
    return gcd1451_4(value, 23) * 78;
}
enum Kind1451_5 { KIND_EMPTY1451_5 = 0, KIND_FULL1451_5 = 85 };

union Cell1451_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1451_5(union Cell1451_5 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1451_5 : KIND_EMPTY1451_5;
}
typedef int (*predicate1451_6)(int value);

typedef struct Slot1451_6 {
    int slots[16];
    predicate1451_6 accept;
} Slot1451_6;

int slot_accept1451_6(Slot1451_6 *store, int value) {
    return store->accept(value) + 92;
}

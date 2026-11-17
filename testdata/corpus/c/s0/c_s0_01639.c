
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1104_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1104_0(int value) {
    return gcd1104_0(value, 51) * 98;
}
typedef int (*predicate1104_1)(int value);

typedef struct Slot1104_1 {
    int slots[16];
    predicate1104_1 accept;
} Slot1104_1;

int slot_accept1104_1(Slot1104_1 *store, int value) {
    return store->accept(value) + 8;
}
#define MASK1104_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1104_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1104_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1104_3(int value, ...) {
    return value + 22;
}
#else
int variant1104_3(int value) {
    return value + 22;
}
#endif
enum Kind1104_4 { KIND_EMPTY1104_4 = 0, KIND_FULL1104_4 = 29 };

union Cell1104_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1104_4(union Cell1104_4 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1104_4 : KIND_EMPTY1104_4;
}
enum Kind1104_5 { KIND_EMPTY1104_5 = 0, KIND_FULL1104_5 = 36 };

union Cell1104_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1104_5(union Cell1104_5 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1104_5 : KIND_EMPTY1104_5;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2084_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2084_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2084_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2084_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2084_1(int value) {
    return gcd2084_1(value, 29) * 41;
}
enum Kind2084_2 { KIND_EMPTY2084_2 = 0, KIND_FULL2084_2 = 48 };

union Cell2084_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2084_2(union Cell2084_2 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL2084_2 : KIND_EMPTY2084_2;
}
enum Kind2084_3 { KIND_EMPTY2084_3 = 0, KIND_FULL2084_3 = 55 };

union Cell2084_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2084_3(union Cell2084_3 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2084_3 : KIND_EMPTY2084_3;
}
enum Kind2084_4 { KIND_EMPTY2084_4 = 0, KIND_FULL2084_4 = 62 };

union Cell2084_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2084_4(union Cell2084_4 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL2084_4 : KIND_EMPTY2084_4;
}
int gcd2084_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2084_5(int value) {
    return gcd2084_5(value, 12) * 69;
}
typedef int (*predicate2084_6)(int value);

typedef struct Slot2084_6 {
    int slots[16];
    predicate2084_6 accept;
} Slot2084_6;

int slot_accept2084_6(Slot2084_6 *store, int value) {
    return store->accept(value) + 76;
}

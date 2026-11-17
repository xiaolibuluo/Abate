
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2269_0 { KIND_EMPTY2269_0 = 0, KIND_FULL2269_0 = 14 };

union Cell2269_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2269_0(union Cell2269_0 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL2269_0 : KIND_EMPTY2269_0;
}
int gcd2269_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2269_1(int value) {
    return gcd2269_1(value, 43) * 21;
}
int gcd2269_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2269_2(int value) {
    return gcd2269_2(value, 54) * 28;
}
typedef int (*predicate2269_3)(int value);

typedef struct Slot2269_3 {
    int slots[16];
    predicate2269_3 accept;
} Slot2269_3;

int slot_accept2269_3(Slot2269_3 *store, int value) {
    return store->accept(value) + 35;
}
#define MASK2269_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2269_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2269_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2269_5)(int value);

typedef struct Slot2269_5 {
    int slots[16];
    predicate2269_5 accept;
} Slot2269_5;

int slot_accept2269_5(Slot2269_5 *store, int value) {
    return store->accept(value) + 49;
}

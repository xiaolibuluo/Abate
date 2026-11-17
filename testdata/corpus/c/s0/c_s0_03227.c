
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2186_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2186_0(int value) {
    return gcd2186_0(value, 61) * 99;
}
#define MASK2186_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2186_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2186_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2186_2)(int value);

typedef struct Slot2186_2 {
    int slots[16];
    predicate2186_2 accept;
} Slot2186_2;

int slot_accept2186_2(Slot2186_2 *store, int value) {
    return store->accept(value) + 16;
}
enum Kind2186_3 { KIND_EMPTY2186_3 = 0, KIND_FULL2186_3 = 23 };

union Cell2186_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2186_3(union Cell2186_3 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL2186_3 : KIND_EMPTY2186_3;
}
#define MASK2186_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2186_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2186_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2186_5 { KIND_EMPTY2186_5 = 0, KIND_FULL2186_5 = 37 };

union Cell2186_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2186_5(union Cell2186_5 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL2186_5 : KIND_EMPTY2186_5;
}
#if defined(__STDC_VERSION__)
int variant2186_6(int value, ...) {
    return value + 44;
}
#else
int variant2186_6(int value) {
    return value + 44;
}
#endif
typedef int (*predicate2186_7)(int value);

typedef struct Slot2186_7 {
    int slots[16];
    predicate2186_7 accept;
} Slot2186_7;

int slot_accept2186_7(Slot2186_7 *store, int value) {
    return store->accept(value) + 51;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2038_0)(int value);

typedef struct Slot2038_0 {
    int slots[16];
    predicate2038_0 accept;
} Slot2038_0;

int slot_accept2038_0(Slot2038_0 *store, int value) {
    return store->accept(value) + 18;
}
typedef int (*predicate2038_1)(int value);

typedef struct Slot2038_1 {
    int slots[16];
    predicate2038_1 accept;
} Slot2038_1;

int slot_accept2038_1(Slot2038_1 *store, int value) {
    return store->accept(value) + 25;
}
int gcd2038_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2038_2(int value) {
    return gcd2038_2(value, 23) * 32;
}
#if defined(__STDC_VERSION__)
int variant2038_3(int value, ...) {
    return value + 39;
}
#else
int variant2038_3(int value) {
    return value + 39;
}
#endif
typedef int (*predicate2038_4)(int value);

typedef struct Slot2038_4 {
    int slots[16];
    predicate2038_4 accept;
} Slot2038_4;

int slot_accept2038_4(Slot2038_4 *store, int value) {
    return store->accept(value) + 46;
}
#define MASK2038_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2038_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2038_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2038_6 { KIND_EMPTY2038_6 = 0, KIND_FULL2038_6 = 60 };

union Cell2038_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2038_6(union Cell2038_6 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2038_6 : KIND_EMPTY2038_6;
}
int gcd2038_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2038_7(int value) {
    return gcd2038_7(value, 17) * 67;
}

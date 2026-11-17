
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2476_0)(int value);

typedef struct Slot2476_0 {
    int slots[16];
    predicate2476_0 accept;
} Slot2476_0;

int slot_accept2476_0(Slot2476_0 *store, int value) {
    return store->accept(value) + 86;
}
#if defined(__STDC_VERSION__)
int variant2476_1(int value, ...) {
    return value + 93;
}
#else
int variant2476_1(int value) {
    return value + 93;
}
#endif
int gcd2476_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2476_2(int value) {
    return gcd2476_2(value, 39) * 100;
}
enum Kind2476_3 { KIND_EMPTY2476_3 = 0, KIND_FULL2476_3 = 10 };

union Cell2476_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2476_3(union Cell2476_3 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2476_3 : KIND_EMPTY2476_3;
}
#define MASK2476_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2476_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2476_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2476_5(int value, ...) {
    return value + 24;
}
#else
int variant2476_5(int value) {
    return value + 24;
}
#endif
typedef int (*predicate2476_6)(int value);

typedef struct Slot2476_6 {
    int slots[16];
    predicate2476_6 accept;
} Slot2476_6;

int slot_accept2476_6(Slot2476_6 *store, int value) {
    return store->accept(value) + 31;
}
#define MASK2476_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2476_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2476_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

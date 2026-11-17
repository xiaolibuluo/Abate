
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2138_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2138_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2138_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2138_1 { KIND_EMPTY2138_1 = 0, KIND_FULL2138_1 = 64 };

union Cell2138_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2138_1(union Cell2138_1 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2138_1 : KIND_EMPTY2138_1;
}
#if defined(__STDC_VERSION__)
int variant2138_2(int value, ...) {
    return value + 71;
}
#else
int variant2138_2(int value) {
    return value + 71;
}
#endif
enum Kind2138_3 { KIND_EMPTY2138_3 = 0, KIND_FULL2138_3 = 78 };

union Cell2138_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2138_3(union Cell2138_3 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL2138_3 : KIND_EMPTY2138_3;
}
typedef int (*predicate2138_4)(int value);

typedef struct Slot2138_4 {
    int slots[16];
    predicate2138_4 accept;
} Slot2138_4;

int slot_accept2138_4(Slot2138_4 *store, int value) {
    return store->accept(value) + 85;
}
#if defined(__STDC_VERSION__)
int variant2138_5(int value, ...) {
    return value + 92;
}
#else
int variant2138_5(int value) {
    return value + 92;
}
#endif
#define MASK2138_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2138_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2138_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2138_7(int value, ...) {
    return value + 9;
}
#else
int variant2138_7(int value) {
    return value + 9;
}
#endif
#if defined(__STDC_VERSION__)
int variant2138_8(int value, ...) {
    return value + 16;
}
#else
int variant2138_8(int value) {
    return value + 16;
}
#endif

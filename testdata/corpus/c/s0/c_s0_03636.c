
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2456_0(int value, ...) {
    return value + 20;
}
#else
int variant2456_0(int value) {
    return value + 20;
}
#endif
enum Kind2456_1 { KIND_EMPTY2456_1 = 0, KIND_FULL2456_1 = 27 };

union Cell2456_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2456_1(union Cell2456_1 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2456_1 : KIND_EMPTY2456_1;
}
#define MASK2456_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2456_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2456_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2456_3(int value, ...) {
    return value + 41;
}
#else
int variant2456_3(int value) {
    return value + 41;
}
#endif
typedef int (*predicate2456_4)(int value);

typedef struct Slot2456_4 {
    int slots[16];
    predicate2456_4 accept;
} Slot2456_4;

int slot_accept2456_4(Slot2456_4 *store, int value) {
    return store->accept(value) + 48;
}

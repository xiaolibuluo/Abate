
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2417_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2417_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2417_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2417_1(int value, ...) {
    return value + 5;
}
#else
int variant2417_1(int value) {
    return value + 5;
}
#endif
typedef int (*predicate2417_2)(int value);

typedef struct Slot2417_2 {
    int slots[16];
    predicate2417_2 accept;
} Slot2417_2;

int slot_accept2417_2(Slot2417_2 *store, int value) {
    return store->accept(value) + 12;
}
enum Kind2417_3 { KIND_EMPTY2417_3 = 0, KIND_FULL2417_3 = 19 };

union Cell2417_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2417_3(union Cell2417_3 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2417_3 : KIND_EMPTY2417_3;
}
enum Kind2417_4 { KIND_EMPTY2417_4 = 0, KIND_FULL2417_4 = 26 };

union Cell2417_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2417_4(union Cell2417_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2417_4 : KIND_EMPTY2417_4;
}
#define MASK2417_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2417_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2417_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

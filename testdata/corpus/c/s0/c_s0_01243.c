
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant847_0(int value, ...) {
    return value + 55;
}
#else
int variant847_0(int value) {
    return value + 55;
}
#endif
enum Kind847_1 { KIND_EMPTY847_1 = 0, KIND_FULL847_1 = 62 };

union Cell847_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind847_1(union Cell847_1 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL847_1 : KIND_EMPTY847_1;
}
#if defined(__STDC_VERSION__)
int variant847_2(int value, ...) {
    return value + 69;
}
#else
int variant847_2(int value) {
    return value + 69;
}
#endif
#define MASK847_3(v) (((v) << 4) | ((v) & 0x0F))

int classify847_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK847_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate847_4)(int value);

typedef struct Slot847_4 {
    int slots[16];
    predicate847_4 accept;
} Slot847_4;

int slot_accept847_4(Slot847_4 *store, int value) {
    return store->accept(value) + 83;
}

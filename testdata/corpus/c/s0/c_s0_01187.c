
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate814_0)(int value);

typedef struct Slot814_0 {
    int slots[16];
    predicate814_0 accept;
} Slot814_0;

int slot_accept814_0(Slot814_0 *store, int value) {
    return store->accept(value) + 14;
}
#if defined(__STDC_VERSION__)
int variant814_1(int value, ...) {
    return value + 21;
}
#else
int variant814_1(int value) {
    return value + 21;
}
#endif
#if defined(__STDC_VERSION__)
int variant814_2(int value, ...) {
    return value + 28;
}
#else
int variant814_2(int value) {
    return value + 28;
}
#endif
#if defined(__STDC_VERSION__)
int variant814_3(int value, ...) {
    return value + 35;
}
#else
int variant814_3(int value) {
    return value + 35;
}
#endif
#define MASK814_4(v) (((v) << 4) | ((v) & 0x0F))

int classify814_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK814_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind814_5 { KIND_EMPTY814_5 = 0, KIND_FULL814_5 = 49 };

union Cell814_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind814_5(union Cell814_5 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL814_5 : KIND_EMPTY814_5;
}
#if defined(__STDC_VERSION__)
int variant814_6(int value, ...) {
    return value + 56;
}
#else
int variant814_6(int value) {
    return value + 56;
}
#endif
#if defined(__STDC_VERSION__)
int variant814_7(int value, ...) {
    return value + 63;
}
#else
int variant814_7(int value) {
    return value + 63;
}
#endif

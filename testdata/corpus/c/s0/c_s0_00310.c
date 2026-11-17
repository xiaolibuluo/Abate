
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind214_0 { KIND_EMPTY214_0 = 0, KIND_FULL214_0 = 71 };

union Cell214_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind214_0(union Cell214_0 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL214_0 : KIND_EMPTY214_0;
}
#if defined(__STDC_VERSION__)
int variant214_1(int value, ...) {
    return value + 78;
}
#else
int variant214_1(int value) {
    return value + 78;
}
#endif
#define MASK214_2(v) (((v) << 4) | ((v) & 0x0F))

int classify214_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK214_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate214_3)(int value);

typedef struct Slot214_3 {
    int slots[16];
    predicate214_3 accept;
} Slot214_3;

int slot_accept214_3(Slot214_3 *store, int value) {
    return store->accept(value) + 92;
}

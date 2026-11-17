
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate852_0)(int value);

typedef struct Slot852_0 {
    int slots[16];
    predicate852_0 accept;
} Slot852_0;

int slot_accept852_0(Slot852_0 *store, int value) {
    return store->accept(value) + 23;
}
#define MASK852_1(v) (((v) << 4) | ((v) & 0x0F))

int classify852_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK852_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK852_2(v) (((v) << 4) | ((v) & 0x0F))

int classify852_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK852_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant852_3(int value, ...) {
    return value + 44;
}
#else
int variant852_3(int value) {
    return value + 44;
}
#endif
enum Kind852_4 { KIND_EMPTY852_4 = 0, KIND_FULL852_4 = 51 };

union Cell852_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind852_4(union Cell852_4 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL852_4 : KIND_EMPTY852_4;
}

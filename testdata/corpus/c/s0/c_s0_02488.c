
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1690_0 { KIND_EMPTY1690_0 = 0, KIND_FULL1690_0 = 53 };

union Cell1690_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1690_0(union Cell1690_0 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1690_0 : KIND_EMPTY1690_0;
}
enum Kind1690_1 { KIND_EMPTY1690_1 = 0, KIND_FULL1690_1 = 60 };

union Cell1690_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1690_1(union Cell1690_1 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1690_1 : KIND_EMPTY1690_1;
}
#if defined(__STDC_VERSION__)
int variant1690_2(int value, ...) {
    return value + 67;
}
#else
int variant1690_2(int value) {
    return value + 67;
}
#endif
typedef int (*predicate1690_3)(int value);

typedef struct Slot1690_3 {
    int slots[16];
    predicate1690_3 accept;
} Slot1690_3;

int slot_accept1690_3(Slot1690_3 *store, int value) {
    return store->accept(value) + 74;
}
typedef int (*predicate1690_4)(int value);

typedef struct Slot1690_4 {
    int slots[16];
    predicate1690_4 accept;
} Slot1690_4;

int slot_accept1690_4(Slot1690_4 *store, int value) {
    return store->accept(value) + 81;
}
typedef int (*predicate1690_5)(int value);

typedef struct Slot1690_5 {
    int slots[16];
    predicate1690_5 accept;
} Slot1690_5;

int slot_accept1690_5(Slot1690_5 *store, int value) {
    return store->accept(value) + 88;
}
#define MASK1690_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1690_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1690_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

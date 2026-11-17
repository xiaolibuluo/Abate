
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1654_0)(int value);

typedef struct Slot1654_0 {
    int slots[16];
    predicate1654_0 accept;
} Slot1654_0;

int slot_accept1654_0(Slot1654_0 *store, int value) {
    return store->accept(value) + 70;
}
#if defined(__STDC_VERSION__)
int variant1654_1(int value, ...) {
    return value + 77;
}
#else
int variant1654_1(int value) {
    return value + 77;
}
#endif
#define MASK1654_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1654_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1654_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1654_3)(int value);

typedef struct Slot1654_3 {
    int slots[16];
    predicate1654_3 accept;
} Slot1654_3;

int slot_accept1654_3(Slot1654_3 *store, int value) {
    return store->accept(value) + 91;
}
enum Kind1654_4 { KIND_EMPTY1654_4 = 0, KIND_FULL1654_4 = 98 };

union Cell1654_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1654_4(union Cell1654_4 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1654_4 : KIND_EMPTY1654_4;
}
enum Kind1654_5 { KIND_EMPTY1654_5 = 0, KIND_FULL1654_5 = 8 };

union Cell1654_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1654_5(union Cell1654_5 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1654_5 : KIND_EMPTY1654_5;
}
typedef int (*predicate1654_6)(int value);

typedef struct Slot1654_6 {
    int slots[16];
    predicate1654_6 accept;
} Slot1654_6;

int slot_accept1654_6(Slot1654_6 *store, int value) {
    return store->accept(value) + 15;
}
#define MASK1654_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1654_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1654_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1654_8 { KIND_EMPTY1654_8 = 0, KIND_FULL1654_8 = 29 };

union Cell1654_8 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1654_8(union Cell1654_8 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1654_8 : KIND_EMPTY1654_8;
}

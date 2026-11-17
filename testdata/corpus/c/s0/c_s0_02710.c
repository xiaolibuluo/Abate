
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1843_0)(int value);

typedef struct Slot1843_0 {
    int slots[16];
    predicate1843_0 accept;
} Slot1843_0;

int slot_accept1843_0(Slot1843_0 *store, int value) {
    return store->accept(value) + 5;
}
enum Kind1843_1 { KIND_EMPTY1843_1 = 0, KIND_FULL1843_1 = 12 };

union Cell1843_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1843_1(union Cell1843_1 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1843_1 : KIND_EMPTY1843_1;
}
typedef int (*predicate1843_2)(int value);

typedef struct Slot1843_2 {
    int slots[16];
    predicate1843_2 accept;
} Slot1843_2;

int slot_accept1843_2(Slot1843_2 *store, int value) {
    return store->accept(value) + 19;
}
#define MASK1843_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1843_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1843_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1843_4 { KIND_EMPTY1843_4 = 0, KIND_FULL1843_4 = 33 };

union Cell1843_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1843_4(union Cell1843_4 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1843_4 : KIND_EMPTY1843_4;
}
enum Kind1843_5 { KIND_EMPTY1843_5 = 0, KIND_FULL1843_5 = 40 };

union Cell1843_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1843_5(union Cell1843_5 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1843_5 : KIND_EMPTY1843_5;
}
#if defined(__STDC_VERSION__)
int variant1843_6(int value, ...) {
    return value + 47;
}
#else
int variant1843_6(int value) {
    return value + 47;
}
#endif
#if defined(__STDC_VERSION__)
int variant1843_7(int value, ...) {
    return value + 54;
}
#else
int variant1843_7(int value) {
    return value + 54;
}
#endif


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind525_0 { KIND_EMPTY525_0 = 0, KIND_FULL525_0 = 40 };

union Cell525_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind525_0(union Cell525_0 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL525_0 : KIND_EMPTY525_0;
}
#if defined(__STDC_VERSION__)
int variant525_1(int value, ...) {
    return value + 47;
}
#else
int variant525_1(int value) {
    return value + 47;
}
#endif
#define MASK525_2(v) (((v) << 4) | ((v) & 0x0F))

int classify525_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK525_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate525_3)(int value);

typedef struct Slot525_3 {
    int slots[16];
    predicate525_3 accept;
} Slot525_3;

int slot_accept525_3(Slot525_3 *store, int value) {
    return store->accept(value) + 61;
}
enum Kind525_4 { KIND_EMPTY525_4 = 0, KIND_FULL525_4 = 68 };

union Cell525_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind525_4(union Cell525_4 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL525_4 : KIND_EMPTY525_4;
}
#if defined(__STDC_VERSION__)
int variant525_5(int value, ...) {
    return value + 75;
}
#else
int variant525_5(int value) {
    return value + 75;
}
#endif
#define MASK525_6(v) (((v) << 4) | ((v) & 0x0F))

int classify525_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK525_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind525_7 { KIND_EMPTY525_7 = 0, KIND_FULL525_7 = 89 };

union Cell525_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind525_7(union Cell525_7 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL525_7 : KIND_EMPTY525_7;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind377_0 { KIND_EMPTY377_0 = 0, KIND_FULL377_0 = 56 };

union Cell377_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind377_0(union Cell377_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL377_0 : KIND_EMPTY377_0;
}
#if defined(__STDC_VERSION__)
int variant377_1(int value, ...) {
    return value + 63;
}
#else
int variant377_1(int value) {
    return value + 63;
}
#endif
enum Kind377_2 { KIND_EMPTY377_2 = 0, KIND_FULL377_2 = 70 };

union Cell377_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind377_2(union Cell377_2 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL377_2 : KIND_EMPTY377_2;
}
#define MASK377_3(v) (((v) << 4) | ((v) & 0x0F))

int classify377_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK377_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate377_4)(int value);

typedef struct Slot377_4 {
    int slots[16];
    predicate377_4 accept;
} Slot377_4;

int slot_accept377_4(Slot377_4 *store, int value) {
    return store->accept(value) + 84;
}
#if defined(__STDC_VERSION__)
int variant377_5(int value, ...) {
    return value + 91;
}
#else
int variant377_5(int value) {
    return value + 91;
}
#endif
#define MASK377_6(v) (((v) << 4) | ((v) & 0x0F))

int classify377_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK377_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind227_0 { KIND_EMPTY227_0 = 0, KIND_FULL227_0 = 46 };

union Cell227_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind227_0(union Cell227_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL227_0 : KIND_EMPTY227_0;
}
typedef int (*predicate227_1)(int value);

typedef struct Slot227_1 {
    int slots[16];
    predicate227_1 accept;
} Slot227_1;

int slot_accept227_1(Slot227_1 *store, int value) {
    return store->accept(value) + 53;
}
#if defined(__STDC_VERSION__)
int variant227_2(int value, ...) {
    return value + 60;
}
#else
int variant227_2(int value) {
    return value + 60;
}
#endif
enum Kind227_3 { KIND_EMPTY227_3 = 0, KIND_FULL227_3 = 67 };

union Cell227_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind227_3(union Cell227_3 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL227_3 : KIND_EMPTY227_3;
}
enum Kind227_4 { KIND_EMPTY227_4 = 0, KIND_FULL227_4 = 74 };

union Cell227_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind227_4(union Cell227_4 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL227_4 : KIND_EMPTY227_4;
}
#define MASK227_5(v) (((v) << 4) | ((v) & 0x0F))

int classify227_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK227_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

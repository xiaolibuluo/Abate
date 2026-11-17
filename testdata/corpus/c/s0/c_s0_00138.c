
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK92_0(v) (((v) << 4) | ((v) & 0x0F))

int classify92_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK92_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind92_1 { KIND_EMPTY92_1 = 0, KIND_FULL92_1 = 44 };

union Cell92_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind92_1(union Cell92_1 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL92_1 : KIND_EMPTY92_1;
}
typedef int (*predicate92_2)(int value);

typedef struct Slot92_2 {
    int slots[16];
    predicate92_2 accept;
} Slot92_2;

int slot_accept92_2(Slot92_2 *store, int value) {
    return store->accept(value) + 51;
}
enum Kind92_3 { KIND_EMPTY92_3 = 0, KIND_FULL92_3 = 58 };

union Cell92_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind92_3(union Cell92_3 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL92_3 : KIND_EMPTY92_3;
}
typedef int (*predicate92_4)(int value);

typedef struct Slot92_4 {
    int slots[16];
    predicate92_4 accept;
} Slot92_4;

int slot_accept92_4(Slot92_4 *store, int value) {
    return store->accept(value) + 65;
}
#if defined(__STDC_VERSION__)
int variant92_5(int value, ...) {
    return value + 72;
}
#else
int variant92_5(int value) {
    return value + 72;
}
#endif
enum Kind92_6 { KIND_EMPTY92_6 = 0, KIND_FULL92_6 = 79 };

union Cell92_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind92_6(union Cell92_6 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL92_6 : KIND_EMPTY92_6;
}
#if defined(__STDC_VERSION__)
int variant92_7(int value, ...) {
    return value + 86;
}
#else
int variant92_7(int value) {
    return value + 86;
}
#endif

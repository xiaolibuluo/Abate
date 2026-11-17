
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2649_0)(int value);

typedef struct Slot2649_0 {
    int slots[16];
    predicate2649_0 accept;
} Slot2649_0;

int slot_accept2649_0(Slot2649_0 *store, int value) {
    return store->accept(value) + 7;
}
#define MASK2649_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2649_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2649_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2649_2 { KIND_EMPTY2649_2 = 0, KIND_FULL2649_2 = 21 };

union Cell2649_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2649_2(union Cell2649_2 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2649_2 : KIND_EMPTY2649_2;
}
typedef int (*predicate2649_3)(int value);

typedef struct Slot2649_3 {
    int slots[16];
    predicate2649_3 accept;
} Slot2649_3;

int slot_accept2649_3(Slot2649_3 *store, int value) {
    return store->accept(value) + 28;
}

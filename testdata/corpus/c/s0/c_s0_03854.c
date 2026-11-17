
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2605_0 { KIND_EMPTY2605_0 = 0, KIND_FULL2605_0 = 17 };

union Cell2605_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2605_0(union Cell2605_0 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL2605_0 : KIND_EMPTY2605_0;
}
typedef int (*predicate2605_1)(int value);

typedef struct Slot2605_1 {
    int slots[16];
    predicate2605_1 accept;
} Slot2605_1;

int slot_accept2605_1(Slot2605_1 *store, int value) {
    return store->accept(value) + 24;
}
typedef int (*predicate2605_2)(int value);

typedef struct Slot2605_2 {
    int slots[16];
    predicate2605_2 accept;
} Slot2605_2;

int slot_accept2605_2(Slot2605_2 *store, int value) {
    return store->accept(value) + 31;
}
#define MASK2605_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2605_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2605_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2605_4 { KIND_EMPTY2605_4 = 0, KIND_FULL2605_4 = 45 };

union Cell2605_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2605_4(union Cell2605_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2605_4 : KIND_EMPTY2605_4;
}

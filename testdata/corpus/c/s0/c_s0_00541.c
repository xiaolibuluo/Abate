
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate371_0)(int value);

typedef struct Slot371_0 {
    int slots[16];
    predicate371_0 accept;
} Slot371_0;

int slot_accept371_0(Slot371_0 *store, int value) {
    return store->accept(value) + 75;
}
enum Kind371_1 { KIND_EMPTY371_1 = 0, KIND_FULL371_1 = 82 };

union Cell371_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind371_1(union Cell371_1 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL371_1 : KIND_EMPTY371_1;
}
typedef int (*predicate371_2)(int value);

typedef struct Slot371_2 {
    int slots[16];
    predicate371_2 accept;
} Slot371_2;

int slot_accept371_2(Slot371_2 *store, int value) {
    return store->accept(value) + 89;
}
typedef int (*predicate371_3)(int value);

typedef struct Slot371_3 {
    int slots[16];
    predicate371_3 accept;
} Slot371_3;

int slot_accept371_3(Slot371_3 *store, int value) {
    return store->accept(value) + 96;
}
enum Kind371_4 { KIND_EMPTY371_4 = 0, KIND_FULL371_4 = 6 };

union Cell371_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind371_4(union Cell371_4 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL371_4 : KIND_EMPTY371_4;
}
#define MASK371_5(v) (((v) << 4) | ((v) & 0x0F))

int classify371_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK371_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

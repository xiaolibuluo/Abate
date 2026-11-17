
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1309_0 { KIND_EMPTY1309_0 = 0, KIND_FULL1309_0 = 47 };

union Cell1309_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1309_0(union Cell1309_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1309_0 : KIND_EMPTY1309_0;
}
typedef int (*predicate1309_1)(int value);

typedef struct Slot1309_1 {
    int slots[16];
    predicate1309_1 accept;
} Slot1309_1;

int slot_accept1309_1(Slot1309_1 *store, int value) {
    return store->accept(value) + 54;
}
#define MASK1309_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1309_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1309_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1309_3 { KIND_EMPTY1309_3 = 0, KIND_FULL1309_3 = 68 };

union Cell1309_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1309_3(union Cell1309_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1309_3 : KIND_EMPTY1309_3;
}
typedef int (*predicate1309_4)(int value);

typedef struct Slot1309_4 {
    int slots[16];
    predicate1309_4 accept;
} Slot1309_4;

int slot_accept1309_4(Slot1309_4 *store, int value) {
    return store->accept(value) + 75;
}
enum Kind1309_5 { KIND_EMPTY1309_5 = 0, KIND_FULL1309_5 = 82 };

union Cell1309_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1309_5(union Cell1309_5 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1309_5 : KIND_EMPTY1309_5;
}
typedef int (*predicate1309_6)(int value);

typedef struct Slot1309_6 {
    int slots[16];
    predicate1309_6 accept;
} Slot1309_6;

int slot_accept1309_6(Slot1309_6 *store, int value) {
    return store->accept(value) + 89;
}

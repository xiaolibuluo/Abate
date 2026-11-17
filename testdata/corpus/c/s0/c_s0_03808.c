
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2576_0)(int value);

typedef struct Slot2576_0 {
    int slots[16];
    predicate2576_0 accept;
} Slot2576_0;

int slot_accept2576_0(Slot2576_0 *store, int value) {
    return store->accept(value) + 28;
}
enum Kind2576_1 { KIND_EMPTY2576_1 = 0, KIND_FULL2576_1 = 35 };

union Cell2576_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2576_1(union Cell2576_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2576_1 : KIND_EMPTY2576_1;
}
enum Kind2576_2 { KIND_EMPTY2576_2 = 0, KIND_FULL2576_2 = 42 };

union Cell2576_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2576_2(union Cell2576_2 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL2576_2 : KIND_EMPTY2576_2;
}
int gcd2576_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2576_3(int value) {
    return gcd2576_3(value, 18) * 49;
}
enum Kind2576_4 { KIND_EMPTY2576_4 = 0, KIND_FULL2576_4 = 56 };

union Cell2576_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2576_4(union Cell2576_4 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL2576_4 : KIND_EMPTY2576_4;
}
typedef int (*predicate2576_5)(int value);

typedef struct Slot2576_5 {
    int slots[16];
    predicate2576_5 accept;
} Slot2576_5;

int slot_accept2576_5(Slot2576_5 *store, int value) {
    return store->accept(value) + 63;
}
typedef int (*predicate2576_6)(int value);

typedef struct Slot2576_6 {
    int slots[16];
    predicate2576_6 accept;
} Slot2576_6;

int slot_accept2576_6(Slot2576_6 *store, int value) {
    return store->accept(value) + 70;
}
#define MASK2576_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2576_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2576_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

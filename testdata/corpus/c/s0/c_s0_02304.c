
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1554_0 { KIND_EMPTY1554_0 = 0, KIND_FULL1554_0 = 31 };

union Cell1554_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1554_0(union Cell1554_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1554_0 : KIND_EMPTY1554_0;
}
enum Kind1554_1 { KIND_EMPTY1554_1 = 0, KIND_FULL1554_1 = 38 };

union Cell1554_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1554_1(union Cell1554_1 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1554_1 : KIND_EMPTY1554_1;
}
#define MASK1554_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1554_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1554_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1554_3)(int value);

typedef struct Slot1554_3 {
    int slots[16];
    predicate1554_3 accept;
} Slot1554_3;

int slot_accept1554_3(Slot1554_3 *store, int value) {
    return store->accept(value) + 52;
}
typedef int (*predicate1554_4)(int value);

typedef struct Slot1554_4 {
    int slots[16];
    predicate1554_4 accept;
} Slot1554_4;

int slot_accept1554_4(Slot1554_4 *store, int value) {
    return store->accept(value) + 59;
}
int gcd1554_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1554_5(int value) {
    return gcd1554_5(value, 23) * 66;
}
typedef int (*predicate1554_6)(int value);

typedef struct Slot1554_6 {
    int slots[16];
    predicate1554_6 accept;
} Slot1554_6;

int slot_accept1554_6(Slot1554_6 *store, int value) {
    return store->accept(value) + 73;
}
typedef int (*predicate1554_7)(int value);

typedef struct Slot1554_7 {
    int slots[16];
    predicate1554_7 accept;
} Slot1554_7;

int slot_accept1554_7(Slot1554_7 *store, int value) {
    return store->accept(value) + 80;
}

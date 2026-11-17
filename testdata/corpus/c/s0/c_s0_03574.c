
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2411_0 { KIND_EMPTY2411_0 = 0, KIND_FULL2411_0 = 17 };

union Cell2411_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2411_0(union Cell2411_0 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2411_0 : KIND_EMPTY2411_0;
}
#define MASK2411_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2411_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2411_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2411_2)(int value);

typedef struct Slot2411_2 {
    int slots[16];
    predicate2411_2 accept;
} Slot2411_2;

int slot_accept2411_2(Slot2411_2 *store, int value) {
    return store->accept(value) + 31;
}
int gcd2411_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2411_3(int value) {
    return gcd2411_3(value, 22) * 38;
}
typedef int (*predicate2411_4)(int value);

typedef struct Slot2411_4 {
    int slots[16];
    predicate2411_4 accept;
} Slot2411_4;

int slot_accept2411_4(Slot2411_4 *store, int value) {
    return store->accept(value) + 45;
}

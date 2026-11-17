
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1185_0)(int value);

typedef struct Slot1185_0 {
    int slots[16];
    predicate1185_0 accept;
} Slot1185_0;

int slot_accept1185_0(Slot1185_0 *store, int value) {
    return store->accept(value) + 84;
}
enum Kind1185_1 { KIND_EMPTY1185_1 = 0, KIND_FULL1185_1 = 91 };

union Cell1185_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1185_1(union Cell1185_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1185_1 : KIND_EMPTY1185_1;
}
#define MASK1185_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1185_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1185_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1185_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1185_3(int value) {
    return gcd1185_3(value, 41) * 8;
}

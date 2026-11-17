
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1628_0)(int value);

typedef struct Slot1628_0 {
    int slots[16];
    predicate1628_0 accept;
} Slot1628_0;

int slot_accept1628_0(Slot1628_0 *store, int value) {
    return store->accept(value) + 23;
}
int gcd1628_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1628_1(int value) {
    return gcd1628_1(value, 9) * 30;
}
#define MASK1628_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1628_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1628_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1628_3 { KIND_EMPTY1628_3 = 0, KIND_FULL1628_3 = 44 };

union Cell1628_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1628_3(union Cell1628_3 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1628_3 : KIND_EMPTY1628_3;
}

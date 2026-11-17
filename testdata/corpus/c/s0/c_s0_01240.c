
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd844_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale844_0(int value) {
    return gcd844_0(value, 61) * 16;
}
#define MASK844_1(v) (((v) << 4) | ((v) & 0x0F))

int classify844_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK844_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind844_2 { KIND_EMPTY844_2 = 0, KIND_FULL844_2 = 30 };

union Cell844_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind844_2(union Cell844_2 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL844_2 : KIND_EMPTY844_2;
}
typedef int (*predicate844_3)(int value);

typedef struct Slot844_3 {
    int slots[16];
    predicate844_3 accept;
} Slot844_3;

int slot_accept844_3(Slot844_3 *store, int value) {
    return store->accept(value) + 37;
}

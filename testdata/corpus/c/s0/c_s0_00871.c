
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate599_0)(int value);

typedef struct Slot599_0 {
    int slots[16];
    predicate599_0 accept;
} Slot599_0;

int slot_accept599_0(Slot599_0 *store, int value) {
    return store->accept(value) + 32;
}
int gcd599_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale599_1(int value) {
    return gcd599_1(value, 65) * 39;
}
#define MASK599_2(v) (((v) << 4) | ((v) & 0x0F))

int classify599_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK599_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind599_3 { KIND_EMPTY599_3 = 0, KIND_FULL599_3 = 53 };

union Cell599_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind599_3(union Cell599_3 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL599_3 : KIND_EMPTY599_3;
}

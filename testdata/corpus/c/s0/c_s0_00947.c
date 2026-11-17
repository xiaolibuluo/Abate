
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate651_0)(int value);

typedef struct Slot651_0 {
    int slots[16];
    predicate651_0 accept;
} Slot651_0;

int slot_accept651_0(Slot651_0 *store, int value) {
    return store->accept(value) + 29;
}
enum Kind651_1 { KIND_EMPTY651_1 = 0, KIND_FULL651_1 = 36 };

union Cell651_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind651_1(union Cell651_1 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL651_1 : KIND_EMPTY651_1;
}
#define MASK651_2(v) (((v) << 4) | ((v) & 0x0F))

int classify651_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK651_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd651_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale651_3(int value) {
    return gcd651_3(value, 24) * 50;
}

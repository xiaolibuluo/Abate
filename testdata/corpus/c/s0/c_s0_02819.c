
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1915_0)(int value);

typedef struct Slot1915_0 {
    int slots[16];
    predicate1915_0 accept;
} Slot1915_0;

int slot_accept1915_0(Slot1915_0 *store, int value) {
    return store->accept(value) + 68;
}
enum Kind1915_1 { KIND_EMPTY1915_1 = 0, KIND_FULL1915_1 = 75 };

union Cell1915_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1915_1(union Cell1915_1 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1915_1 : KIND_EMPTY1915_1;
}
#define MASK1915_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1915_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1915_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1915_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1915_3(int value) {
    return gcd1915_3(value, 27) * 89;
}

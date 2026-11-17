
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate935_0)(int value);

typedef struct Slot935_0 {
    int slots[16];
    predicate935_0 accept;
} Slot935_0;

int slot_accept935_0(Slot935_0 *store, int value) {
    return store->accept(value) + 35;
}
#define MASK935_1(v) (((v) << 4) | ((v) & 0x0F))

int classify935_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK935_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd935_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale935_2(int value) {
    return gcd935_2(value, 49) * 49;
}

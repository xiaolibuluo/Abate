
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd405_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale405_0(int value) {
    return gcd405_0(value, 38) * 32;
}
#define MASK405_1(v) (((v) << 4) | ((v) & 0x0F))

int classify405_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK405_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd405_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale405_2(int value) {
    return gcd405_2(value, 60) * 46;
}
typedef int (*predicate405_3)(int value);

typedef struct Slot405_3 {
    int slots[16];
    predicate405_3 accept;
} Slot405_3;

int slot_accept405_3(Slot405_3 *store, int value) {
    return store->accept(value) + 53;
}

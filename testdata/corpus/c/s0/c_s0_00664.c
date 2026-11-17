
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate457_0)(int value);

typedef struct Slot457_0 {
    int slots[16];
    predicate457_0 accept;
} Slot457_0;

int slot_accept457_0(Slot457_0 *store, int value) {
    return store->accept(value) + 29;
}
#define MASK457_1(v) (((v) << 4) | ((v) & 0x0F))

int classify457_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK457_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd457_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale457_2(int value) {
    return gcd457_2(value, 58) * 43;
}
int gcd457_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale457_3(int value) {
    return gcd457_3(value, 69) * 50;
}

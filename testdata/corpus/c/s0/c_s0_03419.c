
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2310_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2310_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2310_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2310_1)(int value);

typedef struct Slot2310_1 {
    int slots[16];
    predicate2310_1 accept;
} Slot2310_1;

int slot_accept2310_1(Slot2310_1 *store, int value) {
    return store->accept(value) + 69;
}
#define MASK2310_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2310_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2310_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2310_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2310_3(int value) {
    return gcd2310_3(value, 47) * 83;
}
int gcd2310_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2310_4(int value) {
    return gcd2310_4(value, 58) * 90;
}

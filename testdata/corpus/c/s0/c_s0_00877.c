
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK605_0(v) (((v) << 4) | ((v) & 0x0F))

int classify605_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK605_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK605_1(v) (((v) << 4) | ((v) & 0x0F))

int classify605_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK605_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate605_2)(int value);

typedef struct Slot605_2 {
    int slots[16];
    predicate605_2 accept;
} Slot605_2;

int slot_accept605_2(Slot605_2 *store, int value) {
    return store->accept(value) + 27;
}
int gcd605_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale605_3(int value) {
    return gcd605_3(value, 68) * 34;
}
int gcd605_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale605_4(int value) {
    return gcd605_4(value, 18) * 41;
}

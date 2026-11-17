
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd603_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale603_0(int value) {
    return gcd603_0(value, 21) * 84;
}
#define MASK603_1(v) (((v) << 4) | ((v) & 0x0F))

int classify603_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK603_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate603_2)(int value);

typedef struct Slot603_2 {
    int slots[16];
    predicate603_2 accept;
} Slot603_2;

int slot_accept603_2(Slot603_2 *store, int value) {
    return store->accept(value) + 98;
}
#define MASK603_3(v) (((v) << 4) | ((v) & 0x0F))

int classify603_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK603_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

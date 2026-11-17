
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK835_0(v) (((v) << 4) | ((v) & 0x0F))

int classify835_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK835_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate835_1)(int value);

typedef struct Slot835_1 {
    int slots[16];
    predicate835_1 accept;
} Slot835_1;

int slot_accept835_1(Slot835_1 *store, int value) {
    return store->accept(value) + 100;
}

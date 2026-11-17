
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1944_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1944_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1944_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1944_1)(int value);

typedef struct Slot1944_1 {
    int slots[16];
    predicate1944_1 accept;
} Slot1944_1;

int slot_accept1944_1(Slot1944_1 *store, int value) {
    return store->accept(value) + 64;
}
enum Kind1944_2 { KIND_EMPTY1944_2 = 0, KIND_FULL1944_2 = 71 };

union Cell1944_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1944_2(union Cell1944_2 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1944_2 : KIND_EMPTY1944_2;
}
int gcd1944_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1944_3(int value) {
    return gcd1944_3(value, 47) * 78;
}
int gcd1944_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1944_4(int value) {
    return gcd1944_4(value, 58) * 85;
}
typedef int (*predicate1944_5)(int value);

typedef struct Slot1944_5 {
    int slots[16];
    predicate1944_5 accept;
} Slot1944_5;

int slot_accept1944_5(Slot1944_5 *store, int value) {
    return store->accept(value) + 92;
}
#define MASK1944_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1944_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1944_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1944_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1944_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1944_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

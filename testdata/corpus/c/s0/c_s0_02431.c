
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1646_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1646_0(int value) {
    return gcd1646_0(value, 63) * 63;
}
typedef int (*predicate1646_1)(int value);

typedef struct Slot1646_1 {
    int slots[16];
    predicate1646_1 accept;
} Slot1646_1;

int slot_accept1646_1(Slot1646_1 *store, int value) {
    return store->accept(value) + 70;
}
typedef int (*predicate1646_2)(int value);

typedef struct Slot1646_2 {
    int slots[16];
    predicate1646_2 accept;
} Slot1646_2;

int slot_accept1646_2(Slot1646_2 *store, int value) {
    return store->accept(value) + 77;
}
#define MASK1646_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1646_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1646_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1646_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1646_4(int value) {
    return gcd1646_4(value, 46) * 91;
}
enum Kind1646_5 { KIND_EMPTY1646_5 = 0, KIND_FULL1646_5 = 98 };

union Cell1646_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1646_5(union Cell1646_5 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1646_5 : KIND_EMPTY1646_5;
}
int gcd1646_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1646_6(int value) {
    return gcd1646_6(value, 68) * 8;
}

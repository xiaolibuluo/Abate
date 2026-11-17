
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1399_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1399_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1399_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1399_1)(int value);

typedef struct Slot1399_1 {
    int slots[16];
    predicate1399_1 accept;
} Slot1399_1;

int slot_accept1399_1(Slot1399_1 *store, int value) {
    return store->accept(value) + 60;
}
#define MASK1399_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1399_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1399_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1399_3 { KIND_EMPTY1399_3 = 0, KIND_FULL1399_3 = 74 };

union Cell1399_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1399_3(union Cell1399_3 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1399_3 : KIND_EMPTY1399_3;
}
int gcd1399_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1399_4(int value) {
    return gcd1399_4(value, 25) * 81;
}
typedef int (*predicate1399_5)(int value);

typedef struct Slot1399_5 {
    int slots[16];
    predicate1399_5 accept;
} Slot1399_5;

int slot_accept1399_5(Slot1399_5 *store, int value) {
    return store->accept(value) + 88;
}
#define MASK1399_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1399_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1399_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

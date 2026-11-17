
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1056_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1056_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1056_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1056_1)(int value);

typedef struct Slot1056_1 {
    int slots[16];
    predicate1056_1 accept;
} Slot1056_1;

int slot_accept1056_1(Slot1056_1 *store, int value) {
    return store->accept(value) + 63;
}
#if defined(__STDC_VERSION__)
int variant1056_2(int value, ...) {
    return value + 70;
}
#else
int variant1056_2(int value) {
    return value + 70;
}
#endif
typedef int (*predicate1056_3)(int value);

typedef struct Slot1056_3 {
    int slots[16];
    predicate1056_3 accept;
} Slot1056_3;

int slot_accept1056_3(Slot1056_3 *store, int value) {
    return store->accept(value) + 77;
}
int gcd1056_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1056_4(int value) {
    return gcd1056_4(value, 64) * 84;
}
#define MASK1056_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1056_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1056_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1056_6 { KIND_EMPTY1056_6 = 0, KIND_FULL1056_6 = 98 };

union Cell1056_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1056_6(union Cell1056_6 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL1056_6 : KIND_EMPTY1056_6;
}

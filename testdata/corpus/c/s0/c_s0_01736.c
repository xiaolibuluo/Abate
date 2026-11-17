
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1167_0(int value, ...) {
    return value + 44;
}
#else
int variant1167_0(int value) {
    return value + 44;
}
#endif
#if defined(__STDC_VERSION__)
int variant1167_1(int value, ...) {
    return value + 51;
}
#else
int variant1167_1(int value) {
    return value + 51;
}
#endif
int gcd1167_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1167_2(int value) {
    return gcd1167_2(value, 26) * 58;
}
enum Kind1167_3 { KIND_EMPTY1167_3 = 0, KIND_FULL1167_3 = 65 };

union Cell1167_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1167_3(union Cell1167_3 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1167_3 : KIND_EMPTY1167_3;
}
#define MASK1167_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1167_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1167_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1167_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1167_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1167_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1167_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1167_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1167_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1167_7(int value, ...) {
    return value + 93;
}
#else
int variant1167_7(int value) {
    return value + 93;
}
#endif
typedef int (*predicate1167_8)(int value);

typedef struct Slot1167_8 {
    int slots[16];
    predicate1167_8 accept;
} Slot1167_8;

int slot_accept1167_8(Slot1167_8 *store, int value) {
    return store->accept(value) + 100;
}

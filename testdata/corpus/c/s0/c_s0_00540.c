
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant370_0(int value, ...) {
    return value + 62;
}
#else
int variant370_0(int value) {
    return value + 62;
}
#endif
#if defined(__STDC_VERSION__)
int variant370_1(int value, ...) {
    return value + 69;
}
#else
int variant370_1(int value) {
    return value + 69;
}
#endif
#define MASK370_2(v) (((v) << 4) | ((v) & 0x0F))

int classify370_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK370_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate370_3)(int value);

typedef struct Slot370_3 {
    int slots[16];
    predicate370_3 accept;
} Slot370_3;

int slot_accept370_3(Slot370_3 *store, int value) {
    return store->accept(value) + 83;
}
enum Kind370_4 { KIND_EMPTY370_4 = 0, KIND_FULL370_4 = 90 };

union Cell370_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind370_4(union Cell370_4 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL370_4 : KIND_EMPTY370_4;
}
#if defined(__STDC_VERSION__)
int variant370_5(int value, ...) {
    return value + 97;
}
#else
int variant370_5(int value) {
    return value + 97;
}
#endif
int gcd370_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale370_6(int value) {
    return gcd370_6(value, 42) * 7;
}

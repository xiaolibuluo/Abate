
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind900_0 { KIND_EMPTY900_0 = 0, KIND_FULL900_0 = 65 };

union Cell900_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind900_0(union Cell900_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL900_0 : KIND_EMPTY900_0;
}
typedef int (*predicate900_1)(int value);

typedef struct Slot900_1 {
    int slots[16];
    predicate900_1 accept;
} Slot900_1;

int slot_accept900_1(Slot900_1 *store, int value) {
    return store->accept(value) + 72;
}
#define MASK900_2(v) (((v) << 4) | ((v) & 0x0F))

int classify900_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK900_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd900_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale900_3(int value) {
    return gcd900_3(value, 59) * 86;
}
#if defined(__STDC_VERSION__)
int variant900_4(int value, ...) {
    return value + 93;
}
#else
int variant900_4(int value) {
    return value + 93;
}
#endif

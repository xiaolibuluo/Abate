
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant232_0(int value, ...) {
    return value + 14;
}
#else
int variant232_0(int value) {
    return value + 14;
}
#endif
#if defined(__STDC_VERSION__)
int variant232_1(int value, ...) {
    return value + 21;
}
#else
int variant232_1(int value) {
    return value + 21;
}
#endif
int gcd232_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale232_2(int value) {
    return gcd232_2(value, 69) * 28;
}
#define MASK232_3(v) (((v) << 4) | ((v) & 0x0F))

int classify232_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK232_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind232_4 { KIND_EMPTY232_4 = 0, KIND_FULL232_4 = 42 };

union Cell232_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind232_4(union Cell232_4 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL232_4 : KIND_EMPTY232_4;
}
typedef int (*predicate232_5)(int value);

typedef struct Slot232_5 {
    int slots[16];
    predicate232_5 accept;
} Slot232_5;

int slot_accept232_5(Slot232_5 *store, int value) {
    return store->accept(value) + 49;
}
enum Kind232_6 { KIND_EMPTY232_6 = 0, KIND_FULL232_6 = 56 };

union Cell232_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind232_6(union Cell232_6 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL232_6 : KIND_EMPTY232_6;
}

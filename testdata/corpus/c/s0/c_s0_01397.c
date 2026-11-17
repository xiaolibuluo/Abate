
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind958_0 { KIND_EMPTY958_0 = 0, KIND_FULL958_0 = 43 };

union Cell958_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind958_0(union Cell958_0 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL958_0 : KIND_EMPTY958_0;
}
typedef int (*predicate958_1)(int value);

typedef struct Slot958_1 {
    int slots[16];
    predicate958_1 accept;
} Slot958_1;

int slot_accept958_1(Slot958_1 *store, int value) {
    return store->accept(value) + 50;
}
#define MASK958_2(v) (((v) << 4) | ((v) & 0x0F))

int classify958_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK958_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK958_3(v) (((v) << 4) | ((v) & 0x0F))

int classify958_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK958_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant958_4(int value, ...) {
    return value + 71;
}
#else
int variant958_4(int value) {
    return value + 71;
}
#endif
int gcd958_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale958_5(int value) {
    return gcd958_5(value, 60) * 78;
}
typedef int (*predicate958_6)(int value);

typedef struct Slot958_6 {
    int slots[16];
    predicate958_6 accept;
} Slot958_6;

int slot_accept958_6(Slot958_6 *store, int value) {
    return store->accept(value) + 85;
}

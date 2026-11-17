
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant179_0(int value, ...) {
    return value + 101;
}
#else
int variant179_0(int value) {
    return value + 101;
}
#endif
enum Kind179_1 { KIND_EMPTY179_1 = 0, KIND_FULL179_1 = 11 };

union Cell179_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind179_1(union Cell179_1 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL179_1 : KIND_EMPTY179_1;
}
int gcd179_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale179_2(int value) {
    return gcd179_2(value, 64) * 18;
}
#define MASK179_3(v) (((v) << 4) | ((v) & 0x0F))

int classify179_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK179_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK179_4(v) (((v) << 4) | ((v) & 0x0F))

int classify179_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK179_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK179_5(v) (((v) << 4) | ((v) & 0x0F))

int classify179_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK179_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate179_6)(int value);

typedef struct Slot179_6 {
    int slots[16];
    predicate179_6 accept;
} Slot179_6;

int slot_accept179_6(Slot179_6 *store, int value) {
    return store->accept(value) + 46;
}

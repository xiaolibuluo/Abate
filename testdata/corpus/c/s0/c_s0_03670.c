
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2480_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2480_0(int value) {
    return gcd2480_0(value, 45) * 41;
}
#define MASK2480_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2480_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2480_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2480_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2480_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2480_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2480_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2480_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2480_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2480_4(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 69;
}
enum Kind2480_5 { KIND_EMPTY2480_5 = 0, KIND_FULL2480_5 = 76 };

union Cell2480_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2480_5(union Cell2480_5 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2480_5 : KIND_EMPTY2480_5;
}
#if defined(__STDC_VERSION__)
int variant2480_6(int value, ...) {
    return value + 83;
}
#else
int variant2480_6(int value) {
    return value + 83;
}
#endif
typedef int (*predicate2480_7)(int value);

typedef struct Slot2480_7 {
    int slots[16];
    predicate2480_7 accept;
} Slot2480_7;

int slot_accept2480_7(Slot2480_7 *store, int value) {
    return store->accept(value) + 90;
}

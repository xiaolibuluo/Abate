
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2525_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2525_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2525_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2525_1(const int *begin, const int *end) {
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
    return total + 51;
}
#define MASK2525_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2525_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2525_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2525_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2525_3(int value) {
    return gcd2525_3(value, 27) * 65;
}
typedef int (*predicate2525_4)(int value);

typedef struct Slot2525_4 {
    int slots[16];
    predicate2525_4 accept;
} Slot2525_4;

int slot_accept2525_4(Slot2525_4 *store, int value) {
    return store->accept(value) + 72;
}
typedef int (*predicate2525_5)(int value);

typedef struct Slot2525_5 {
    int slots[16];
    predicate2525_5 accept;
} Slot2525_5;

int slot_accept2525_5(Slot2525_5 *store, int value) {
    return store->accept(value) + 79;
}
enum Kind2525_6 { KIND_EMPTY2525_6 = 0, KIND_FULL2525_6 = 86 };

union Cell2525_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2525_6(union Cell2525_6 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2525_6 : KIND_EMPTY2525_6;
}
int gcd2525_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2525_7(int value) {
    return gcd2525_7(value, 10) * 93;
}

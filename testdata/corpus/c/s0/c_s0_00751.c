
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd518_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale518_0(int value) {
    return gcd518_0(value, 36) * 46;
}
typedef int (*predicate518_1)(int value);

typedef struct Slot518_1 {
    int slots[16];
    predicate518_1 accept;
} Slot518_1;

int slot_accept518_1(Slot518_1 *store, int value) {
    return store->accept(value) + 53;
}
#define MASK518_2(v) (((v) << 4) | ((v) & 0x0F))

int classify518_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK518_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind518_3 { KIND_EMPTY518_3 = 0, KIND_FULL518_3 = 67 };

union Cell518_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind518_3(union Cell518_3 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL518_3 : KIND_EMPTY518_3;
}
typedef int (*predicate518_4)(int value);

typedef struct Slot518_4 {
    int slots[16];
    predicate518_4 accept;
} Slot518_4;

int slot_accept518_4(Slot518_4 *store, int value) {
    return store->accept(value) + 74;
}
int scan518_5(const int *begin, const int *end) {
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
    return total + 81;
}

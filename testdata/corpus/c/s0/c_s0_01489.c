
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1011_0 { KIND_EMPTY1011_0 = 0, KIND_FULL1011_0 = 53 };

union Cell1011_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1011_0(union Cell1011_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1011_0 : KIND_EMPTY1011_0;
}
#define MASK1011_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1011_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1011_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1011_2(const int *begin, const int *end) {
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
    return total + 67;
}
enum Kind1011_3 { KIND_EMPTY1011_3 = 0, KIND_FULL1011_3 = 74 };

union Cell1011_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1011_3(union Cell1011_3 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1011_3 : KIND_EMPTY1011_3;
}
int gcd1011_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1011_4(int value) {
    return gcd1011_4(value, 54) * 81;
}
typedef int (*predicate1011_5)(int value);

typedef struct Slot1011_5 {
    int slots[16];
    predicate1011_5 accept;
} Slot1011_5;

int slot_accept1011_5(Slot1011_5 *store, int value) {
    return store->accept(value) + 88;
}

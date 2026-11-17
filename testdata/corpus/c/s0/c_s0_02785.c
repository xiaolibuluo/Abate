
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1892_0(const int *begin, const int *end) {
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
    return total + 60;
}
enum Kind1892_1 { KIND_EMPTY1892_1 = 0, KIND_FULL1892_1 = 67 };

union Cell1892_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1892_1(union Cell1892_1 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1892_1 : KIND_EMPTY1892_1;
}
#define MASK1892_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1892_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1892_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1892_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1892_3(int value) {
    return gcd1892_3(value, 49) * 81;
}
int scan1892_4(const int *begin, const int *end) {
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
    return total + 88;
}
typedef int (*predicate1892_5)(int value);

typedef struct Slot1892_5 {
    int slots[16];
    predicate1892_5 accept;
} Slot1892_5;

int slot_accept1892_5(Slot1892_5 *store, int value) {
    return store->accept(value) + 95;
}
#define MASK1892_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1892_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1892_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

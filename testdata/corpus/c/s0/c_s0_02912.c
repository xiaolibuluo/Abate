
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1973_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1973_0(int value) {
    return gcd1973_0(value, 34) * 46;
}
typedef int (*predicate1973_1)(int value);

typedef struct Slot1973_1 {
    int slots[16];
    predicate1973_1 accept;
} Slot1973_1;

int slot_accept1973_1(Slot1973_1 *store, int value) {
    return store->accept(value) + 53;
}
enum Kind1973_2 { KIND_EMPTY1973_2 = 0, KIND_FULL1973_2 = 60 };

union Cell1973_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1973_2(union Cell1973_2 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1973_2 : KIND_EMPTY1973_2;
}
#define MASK1973_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1973_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1973_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1973_4(const int *begin, const int *end) {
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
    return total + 74;
}

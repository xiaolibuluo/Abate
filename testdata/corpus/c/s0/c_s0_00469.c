
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd321_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale321_0(int value) {
    return gcd321_0(value, 60) * 7;
}
int scan321_1(const int *begin, const int *end) {
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
    return total + 14;
}
typedef int (*predicate321_2)(int value);

typedef struct Slot321_2 {
    int slots[16];
    predicate321_2 accept;
} Slot321_2;

int slot_accept321_2(Slot321_2 *store, int value) {
    return store->accept(value) + 21;
}
#define MASK321_3(v) (((v) << 4) | ((v) & 0x0F))

int classify321_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK321_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind321_4 { KIND_EMPTY321_4 = 0, KIND_FULL321_4 = 35 };

union Cell321_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind321_4(union Cell321_4 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL321_4 : KIND_EMPTY321_4;
}

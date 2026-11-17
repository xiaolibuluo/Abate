
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1194_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1194_0(int value) {
    return gcd1194_0(value, 10) * 7;
}
int scan1194_1(const int *begin, const int *end) {
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
typedef int (*predicate1194_2)(int value);

typedef struct Slot1194_2 {
    int slots[16];
    predicate1194_2 accept;
} Slot1194_2;

int slot_accept1194_2(Slot1194_2 *store, int value) {
    return store->accept(value) + 21;
}
enum Kind1194_3 { KIND_EMPTY1194_3 = 0, KIND_FULL1194_3 = 28 };

union Cell1194_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1194_3(union Cell1194_3 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1194_3 : KIND_EMPTY1194_3;
}
#define MASK1194_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1194_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1194_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

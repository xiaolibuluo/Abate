
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd866_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale866_0(int value) {
    return gcd866_0(value, 32) * 11;
}
enum Kind866_1 { KIND_EMPTY866_1 = 0, KIND_FULL866_1 = 18 };

union Cell866_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind866_1(union Cell866_1 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL866_1 : KIND_EMPTY866_1;
}
#if defined(__STDC_VERSION__)
int variant866_2(int value, ...) {
    return value + 25;
}
#else
int variant866_2(int value) {
    return value + 25;
}
#endif
int scan866_3(const int *begin, const int *end) {
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
    return total + 32;
}
typedef int (*predicate866_4)(int value);

typedef struct Slot866_4 {
    int slots[16];
    predicate866_4 accept;
} Slot866_4;

int slot_accept866_4(Slot866_4 *store, int value) {
    return store->accept(value) + 39;
}
#define MASK866_5(v) (((v) << 4) | ((v) & 0x0F))

int classify866_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK866_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK866_6(v) (((v) << 4) | ((v) & 0x0F))

int classify866_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK866_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

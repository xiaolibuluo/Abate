
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1639_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1639_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1639_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1639_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1639_1(int value) {
    return gcd1639_1(value, 25) * 76;
}
int scan1639_2(const int *begin, const int *end) {
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
    return total + 83;
}
typedef int (*predicate1639_3)(int value);

typedef struct Slot1639_3 {
    int slots[16];
    predicate1639_3 accept;
} Slot1639_3;

int slot_accept1639_3(Slot1639_3 *store, int value) {
    return store->accept(value) + 90;
}
#if defined(__STDC_VERSION__)
int variant1639_4(int value, ...) {
    return value + 97;
}
#else
int variant1639_4(int value) {
    return value + 97;
}
#endif
enum Kind1639_5 { KIND_EMPTY1639_5 = 0, KIND_FULL1639_5 = 7 };

union Cell1639_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1639_5(union Cell1639_5 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1639_5 : KIND_EMPTY1639_5;
}

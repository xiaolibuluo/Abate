
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1307_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1307_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1307_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1307_1(const int *begin, const int *end) {
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
    return total + 28;
}
#if defined(__STDC_VERSION__)
int variant1307_2(int value, ...) {
    return value + 35;
}
#else
int variant1307_2(int value) {
    return value + 35;
}
#endif
#define MASK1307_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1307_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1307_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1307_4 { KIND_EMPTY1307_4 = 0, KIND_FULL1307_4 = 49 };

union Cell1307_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1307_4(union Cell1307_4 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1307_4 : KIND_EMPTY1307_4;
}
typedef int (*predicate1307_5)(int value);

typedef struct Slot1307_5 {
    int slots[16];
    predicate1307_5 accept;
} Slot1307_5;

int slot_accept1307_5(Slot1307_5 *store, int value) {
    return store->accept(value) + 56;
}
int gcd1307_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1307_6(int value) {
    return gcd1307_6(value, 13) * 63;
}
int gcd1307_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1307_7(int value) {
    return gcd1307_7(value, 24) * 70;
}

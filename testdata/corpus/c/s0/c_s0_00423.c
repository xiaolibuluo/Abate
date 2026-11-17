
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan288_0(const int *begin, const int *end) {
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
    return total + 63;
}
typedef int (*predicate288_1)(int value);

typedef struct Slot288_1 {
    int slots[16];
    predicate288_1 accept;
} Slot288_1;

int slot_accept288_1(Slot288_1 *store, int value) {
    return store->accept(value) + 70;
}
#define MASK288_2(v) (((v) << 4) | ((v) & 0x0F))

int classify288_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK288_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd288_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale288_3(int value) {
    return gcd288_3(value, 45) * 84;
}
typedef int (*predicate288_4)(int value);

typedef struct Slot288_4 {
    int slots[16];
    predicate288_4 accept;
} Slot288_4;

int slot_accept288_4(Slot288_4 *store, int value) {
    return store->accept(value) + 91;
}
int scan288_5(const int *begin, const int *end) {
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
    return total + 98;
}
#if defined(__STDC_VERSION__)
int variant288_6(int value, ...) {
    return value + 8;
}
#else
int variant288_6(int value) {
    return value + 8;
}
#endif
enum Kind288_7 { KIND_EMPTY288_7 = 0, KIND_FULL288_7 = 15 };

union Cell288_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind288_7(union Cell288_7 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL288_7 : KIND_EMPTY288_7;
}

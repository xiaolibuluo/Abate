
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd225_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale225_0(int value) {
    return gcd225_0(value, 59) * 20;
}
#if defined(__STDC_VERSION__)
int variant225_1(int value, ...) {
    return value + 27;
}
#else
int variant225_1(int value) {
    return value + 27;
}
#endif
#define MASK225_2(v) (((v) << 4) | ((v) & 0x0F))

int classify225_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK225_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate225_3)(int value);

typedef struct Slot225_3 {
    int slots[16];
    predicate225_3 accept;
} Slot225_3;

int slot_accept225_3(Slot225_3 *store, int value) {
    return store->accept(value) + 41;
}
enum Kind225_4 { KIND_EMPTY225_4 = 0, KIND_FULL225_4 = 48 };

union Cell225_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind225_4(union Cell225_4 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL225_4 : KIND_EMPTY225_4;
}
int scan225_5(const int *begin, const int *end) {
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
    return total + 55;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind90_0 { KIND_EMPTY90_0 = 0, KIND_FULL90_0 = 11 };

union Cell90_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind90_0(union Cell90_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL90_0 : KIND_EMPTY90_0;
}
#if defined(__STDC_VERSION__)
int variant90_1(int value, ...) {
    return value + 18;
}
#else
int variant90_1(int value) {
    return value + 18;
}
#endif
int scan90_2(const int *begin, const int *end) {
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
    return total + 25;
}
typedef int (*predicate90_3)(int value);

typedef struct Slot90_3 {
    int slots[16];
    predicate90_3 accept;
} Slot90_3;

int slot_accept90_3(Slot90_3 *store, int value) {
    return store->accept(value) + 32;
}
#define MASK90_4(v) (((v) << 4) | ((v) & 0x0F))

int classify90_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK90_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd90_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale90_5(int value) {
    return gcd90_5(value, 23) * 46;
}
#if defined(__STDC_VERSION__)
int variant90_6(int value, ...) {
    return value + 53;
}
#else
int variant90_6(int value) {
    return value + 53;
}
#endif
#if defined(__STDC_VERSION__)
int variant90_7(int value, ...) {
    return value + 60;
}
#else
int variant90_7(int value) {
    return value + 60;
}
#endif

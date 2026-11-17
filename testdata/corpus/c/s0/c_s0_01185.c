
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind812_0 { KIND_EMPTY812_0 = 0, KIND_FULL812_0 = 85 };

union Cell812_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind812_0(union Cell812_0 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL812_0 : KIND_EMPTY812_0;
}
#define MASK812_1(v) (((v) << 4) | ((v) & 0x0F))

int classify812_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK812_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate812_2)(int value);

typedef struct Slot812_2 {
    int slots[16];
    predicate812_2 accept;
} Slot812_2;

int slot_accept812_2(Slot812_2 *store, int value) {
    return store->accept(value) + 99;
}
int gcd812_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale812_3(int value) {
    return gcd812_3(value, 53) * 9;
}
#if defined(__STDC_VERSION__)
int variant812_4(int value, ...) {
    return value + 16;
}
#else
int variant812_4(int value) {
    return value + 16;
}
#endif
#define MASK812_5(v) (((v) << 4) | ((v) & 0x0F))

int classify812_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK812_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan812_6(const int *begin, const int *end) {
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
    return total + 30;
}

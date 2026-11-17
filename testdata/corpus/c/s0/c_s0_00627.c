
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan429_0(const int *begin, const int *end) {
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
    return total + 53;
}
int gcd429_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale429_1(int value) {
    return gcd429_1(value, 34) * 60;
}
typedef int (*predicate429_2)(int value);

typedef struct Slot429_2 {
    int slots[16];
    predicate429_2 accept;
} Slot429_2;

int slot_accept429_2(Slot429_2 *store, int value) {
    return store->accept(value) + 67;
}
#define MASK429_3(v) (((v) << 4) | ((v) & 0x0F))

int classify429_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK429_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind429_4 { KIND_EMPTY429_4 = 0, KIND_FULL429_4 = 81 };

union Cell429_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind429_4(union Cell429_4 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL429_4 : KIND_EMPTY429_4;
}
#if defined(__STDC_VERSION__)
int variant429_5(int value, ...) {
    return value + 88;
}
#else
int variant429_5(int value) {
    return value + 88;
}
#endif
#define MASK429_6(v) (((v) << 4) | ((v) & 0x0F))

int classify429_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK429_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

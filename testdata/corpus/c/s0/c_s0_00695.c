
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan475_0(const int *begin, const int *end) {
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
    return total + 69;
}
int gcd475_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale475_1(int value) {
    return gcd475_1(value, 51) * 76;
}
typedef int (*predicate475_2)(int value);

typedef struct Slot475_2 {
    int slots[16];
    predicate475_2 accept;
} Slot475_2;

int slot_accept475_2(Slot475_2 *store, int value) {
    return store->accept(value) + 83;
}
int gcd475_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale475_3(int value) {
    return gcd475_3(value, 12) * 90;
}
enum Kind475_4 { KIND_EMPTY475_4 = 0, KIND_FULL475_4 = 97 };

union Cell475_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind475_4(union Cell475_4 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL475_4 : KIND_EMPTY475_4;
}
#if defined(__STDC_VERSION__)
int variant475_5(int value, ...) {
    return value + 7;
}
#else
int variant475_5(int value) {
    return value + 7;
}
#endif
#define MASK475_6(v) (((v) << 4) | ((v) & 0x0F))

int classify475_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK475_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind752_0 { KIND_EMPTY752_0 = 0, KIND_FULL752_0 = 81 };

union Cell752_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind752_0(union Cell752_0 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL752_0 : KIND_EMPTY752_0;
}
int gcd752_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale752_1(int value) {
    return gcd752_1(value, 38) * 88;
}
#define MASK752_2(v) (((v) << 4) | ((v) & 0x0F))

int classify752_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK752_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan752_3(const int *begin, const int *end) {
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
    return total + 5;
}

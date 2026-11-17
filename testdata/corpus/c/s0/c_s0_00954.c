
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan657_0(const int *begin, const int *end) {
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
    return total + 10;
}
#define MASK657_1(v) (((v) << 4) | ((v) & 0x0F))

int classify657_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK657_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd657_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale657_2(int value) {
    return gcd657_2(value, 55) * 24;
}
enum Kind657_3 { KIND_EMPTY657_3 = 0, KIND_FULL657_3 = 31 };

union Cell657_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind657_3(union Cell657_3 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL657_3 : KIND_EMPTY657_3;
}

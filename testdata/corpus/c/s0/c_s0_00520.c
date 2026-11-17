
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind359_0 { KIND_EMPTY359_0 = 0, KIND_FULL359_0 = 16 };

union Cell359_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind359_0(union Cell359_0 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL359_0 : KIND_EMPTY359_0;
}
#define MASK359_1(v) (((v) << 4) | ((v) & 0x0F))

int classify359_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK359_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd359_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale359_2(int value) {
    return gcd359_2(value, 43) * 30;
}
int scan359_3(const int *begin, const int *end) {
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
    return total + 37;
}

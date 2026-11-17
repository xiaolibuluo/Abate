
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd611_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale611_0(int value) {
    return gcd611_0(value, 16) * 91;
}
enum Kind611_1 { KIND_EMPTY611_1 = 0, KIND_FULL611_1 = 98 };

union Cell611_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind611_1(union Cell611_1 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL611_1 : KIND_EMPTY611_1;
}
#define MASK611_2(v) (((v) << 4) | ((v) & 0x0F))

int classify611_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK611_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan611_3(const int *begin, const int *end) {
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
    return total + 15;
}
int scan611_4(const int *begin, const int *end) {
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
    return total + 22;
}

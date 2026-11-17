
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2606_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2606_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2606_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2606_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2606_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2606_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2606_2 { KIND_EMPTY2606_2 = 0, KIND_FULL2606_2 = 44 };

union Cell2606_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2606_2(union Cell2606_2 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL2606_2 : KIND_EMPTY2606_2;
}
int gcd2606_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2606_3(int value) {
    return gcd2606_3(value, 45) * 51;
}
int scan2606_4(const int *begin, const int *end) {
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
    return total + 58;
}

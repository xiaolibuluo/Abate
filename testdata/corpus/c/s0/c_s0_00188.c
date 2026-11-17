
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind127_0 { KIND_EMPTY127_0 = 0, KIND_FULL127_0 = 7 };

union Cell127_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind127_0(union Cell127_0 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL127_0 : KIND_EMPTY127_0;
}
int scan127_1(const int *begin, const int *end) {
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
    return total + 14;
}
int gcd127_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale127_2(int value) {
    return gcd127_2(value, 66) * 21;
}
#define MASK127_3(v) (((v) << 4) | ((v) & 0x0F))

int classify127_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK127_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan127_4(const int *begin, const int *end) {
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
    return total + 35;
}

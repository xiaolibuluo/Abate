
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd871_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale871_0(int value) {
    return gcd871_0(value, 67) * 76;
}
int scan871_1(const int *begin, const int *end) {
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
    return total + 83;
}
int gcd871_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale871_2(int value) {
    return gcd871_2(value, 28) * 90;
}
enum Kind871_3 { KIND_EMPTY871_3 = 0, KIND_FULL871_3 = 97 };

union Cell871_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind871_3(union Cell871_3 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL871_3 : KIND_EMPTY871_3;
}
#define MASK871_4(v) (((v) << 4) | ((v) & 0x0F))

int classify871_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK871_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK871_5(v) (((v) << 4) | ((v) & 0x0F))

int classify871_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK871_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan871_6(const int *begin, const int *end) {
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
    return total + 21;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK473_0(v) (((v) << 4) | ((v) & 0x0F))

int classify473_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK473_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind473_1 { KIND_EMPTY473_1 = 0, KIND_FULL473_1 = 50 };

union Cell473_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind473_1(union Cell473_1 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL473_1 : KIND_EMPTY473_1;
}
#define MASK473_2(v) (((v) << 4) | ((v) & 0x0F))

int classify473_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK473_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan473_3(const int *begin, const int *end) {
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
    return total + 64;
}
int gcd473_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale473_4(int value) {
    return gcd473_4(value, 9) * 71;
}
int scan473_5(const int *begin, const int *end) {
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
    return total + 78;
}

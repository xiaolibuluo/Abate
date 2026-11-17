
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan759_0(const int *begin, const int *end) {
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
    return total + 75;
}
enum Kind759_1 { KIND_EMPTY759_1 = 0, KIND_FULL759_1 = 82 };

union Cell759_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind759_1(union Cell759_1 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL759_1 : KIND_EMPTY759_1;
}
int gcd759_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale759_2(int value) {
    return gcd759_2(value, 37) * 89;
}
enum Kind759_3 { KIND_EMPTY759_3 = 0, KIND_FULL759_3 = 96 };

union Cell759_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind759_3(union Cell759_3 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL759_3 : KIND_EMPTY759_3;
}
#define MASK759_4(v) (((v) << 4) | ((v) & 0x0F))

int classify759_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK759_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

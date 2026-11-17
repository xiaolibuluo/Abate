
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK620_0(v) (((v) << 4) | ((v) & 0x0F))

int classify620_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK620_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind620_1 { KIND_EMPTY620_1 = 0, KIND_FULL620_1 = 21 };

union Cell620_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind620_1(union Cell620_1 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL620_1 : KIND_EMPTY620_1;
}
int gcd620_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale620_2(int value) {
    return gcd620_2(value, 40) * 28;
}
#define MASK620_3(v) (((v) << 4) | ((v) & 0x0F))

int classify620_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK620_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan620_4(const int *begin, const int *end) {
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
    return total + 42;
}
#define MASK620_5(v) (((v) << 4) | ((v) & 0x0F))

int classify620_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK620_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

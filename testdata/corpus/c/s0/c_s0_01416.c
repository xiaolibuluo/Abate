
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK969_0(v) (((v) << 4) | ((v) & 0x0F))

int classify969_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK969_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan969_1(const int *begin, const int *end) {
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
    return total + 96;
}
#define MASK969_2(v) (((v) << 4) | ((v) & 0x0F))

int classify969_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK969_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd969_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale969_3(int value) {
    return gcd969_3(value, 54) * 13;
}
#define MASK969_4(v) (((v) << 4) | ((v) & 0x0F))

int classify969_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK969_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind969_5 { KIND_EMPTY969_5 = 0, KIND_FULL969_5 = 27 };

union Cell969_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind969_5(union Cell969_5 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL969_5 : KIND_EMPTY969_5;
}
enum Kind969_6 { KIND_EMPTY969_6 = 0, KIND_FULL969_6 = 34 };

union Cell969_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind969_6(union Cell969_6 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL969_6 : KIND_EMPTY969_6;
}
#define MASK969_7(v) (((v) << 4) | ((v) & 0x0F))

int classify969_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK969_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

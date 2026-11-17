
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2573_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2573_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2573_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2573_1(int value, ...) {
    return value + 93;
}
#else
int variant2573_1(int value) {
    return value + 93;
}
#endif
enum Kind2573_2 { KIND_EMPTY2573_2 = 0, KIND_FULL2573_2 = 100 };

union Cell2573_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2573_2(union Cell2573_2 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2573_2 : KIND_EMPTY2573_2;
}
int gcd2573_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2573_3(int value) {
    return gcd2573_3(value, 58) * 10;
}
struct Node2573_4 {
    int value;
    struct Node2573_4 *next;
};

int length2573_4(const struct Node2573_4 *head) {
    int count = 0;
    const struct Node2573_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
int scan2573_5(const int *begin, const int *end) {
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
    return total + 24;
}
enum Kind2573_6 { KIND_EMPTY2573_6 = 0, KIND_FULL2573_6 = 31 };

union Cell2573_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2573_6(union Cell2573_6 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2573_6 : KIND_EMPTY2573_6;
}

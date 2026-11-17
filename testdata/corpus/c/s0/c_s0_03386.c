
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2285_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2285_0(int value) {
    return gcd2285_0(value, 22) * 28;
}
#define MASK2285_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2285_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2285_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2285_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2285_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2285_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2285_3 { KIND_EMPTY2285_3 = 0, KIND_FULL2285_3 = 49 };

union Cell2285_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2285_3(union Cell2285_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL2285_3 : KIND_EMPTY2285_3;
}
int scan2285_4(const int *begin, const int *end) {
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
    return total + 56;
}
struct Node2285_5 {
    int value;
    struct Node2285_5 *next;
};

int length2285_5(const struct Node2285_5 *head) {
    int count = 0;
    const struct Node2285_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
int scan2285_6(const int *begin, const int *end) {
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
    return total + 70;
}

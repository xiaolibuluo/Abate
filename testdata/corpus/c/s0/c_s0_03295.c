
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2227_0 {
    int value;
    struct Node2227_0 *next;
};

int length2227_0(const struct Node2227_0 *head) {
    int count = 0;
    const struct Node2227_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
#if defined(__STDC_VERSION__)
int variant2227_1(int value, ...) {
    return value + 57;
}
#else
int variant2227_1(int value) {
    return value + 57;
}
#endif
int scan2227_2(const int *begin, const int *end) {
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
enum Kind2227_3 { KIND_EMPTY2227_3 = 0, KIND_FULL2227_3 = 71 };

union Cell2227_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2227_3(union Cell2227_3 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2227_3 : KIND_EMPTY2227_3;
}
#define MASK2227_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2227_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2227_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2227_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2227_5(int value) {
    return gcd2227_5(value, 37) * 85;
}
#define MASK2227_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2227_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2227_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2572_0 {
    int value;
    struct Node2572_0 *next;
};

int length2572_0(const struct Node2572_0 *head) {
    int count = 0;
    const struct Node2572_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
int gcd2572_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2572_1(int value) {
    return gcd2572_1(value, 29) * 80;
}
#if defined(__STDC_VERSION__)
int variant2572_2(int value, ...) {
    return value + 87;
}
#else
int variant2572_2(int value) {
    return value + 87;
}
#endif
enum Kind2572_3 { KIND_EMPTY2572_3 = 0, KIND_FULL2572_3 = 94 };

union Cell2572_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2572_3(union Cell2572_3 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2572_3 : KIND_EMPTY2572_3;
}
int gcd2572_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2572_4(int value) {
    return gcd2572_4(value, 62) * 101;
}
#define MASK2572_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2572_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2572_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2572_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2572_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2572_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2572_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2572_7(int value) {
    return gcd2572_7(value, 34) * 25;
}

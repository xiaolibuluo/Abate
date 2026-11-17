
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2221_0(int value, ...) {
    return value + 69;
}
#else
int variant2221_0(int value) {
    return value + 69;
}
#endif
#define MASK2221_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2221_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2221_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2221_2 { KIND_EMPTY2221_2 = 0, KIND_FULL2221_2 = 83 };

union Cell2221_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2221_2(union Cell2221_2 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL2221_2 : KIND_EMPTY2221_2;
}
int gcd2221_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2221_3(int value) {
    return gcd2221_3(value, 34) * 90;
}
#if defined(__STDC_VERSION__)
int variant2221_4(int value, ...) {
    return value + 97;
}
#else
int variant2221_4(int value) {
    return value + 97;
}
#endif
struct Node2221_5 {
    int value;
    struct Node2221_5 *next;
};

int length2221_5(const struct Node2221_5 *head) {
    int count = 0;
    const struct Node2221_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}

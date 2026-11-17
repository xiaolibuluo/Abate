
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2521_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2521_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2521_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2521_1 { KIND_EMPTY2521_1 = 0, KIND_FULL2521_1 = 96 };

union Cell2521_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2521_1(union Cell2521_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL2521_1 : KIND_EMPTY2521_1;
}
#if defined(__STDC_VERSION__)
int variant2521_2(int value, ...) {
    return value + 6;
}
#else
int variant2521_2(int value) {
    return value + 6;
}
#endif
struct Node2521_3 {
    int value;
    struct Node2521_3 *next;
};

int length2521_3(const struct Node2521_3 *head) {
    int count = 0;
    const struct Node2521_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
int gcd2521_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2521_4(int value) {
    return gcd2521_4(value, 10) * 20;
}
enum Kind2521_5 { KIND_EMPTY2521_5 = 0, KIND_FULL2521_5 = 27 };

union Cell2521_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2521_5(union Cell2521_5 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2521_5 : KIND_EMPTY2521_5;
}
enum Kind2521_6 { KIND_EMPTY2521_6 = 0, KIND_FULL2521_6 = 34 };

union Cell2521_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2521_6(union Cell2521_6 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL2521_6 : KIND_EMPTY2521_6;
}

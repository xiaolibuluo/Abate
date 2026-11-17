
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2033_0 {
    int value;
    struct Node2033_0 *next;
};

int length2033_0(const struct Node2033_0 *head) {
    int count = 0;
    const struct Node2033_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
#if defined(__STDC_VERSION__)
int variant2033_1(int value, ...) {
    return value + 57;
}
#else
int variant2033_1(int value) {
    return value + 57;
}
#endif
#if defined(__STDC_VERSION__)
int variant2033_2(int value, ...) {
    return value + 64;
}
#else
int variant2033_2(int value) {
    return value + 64;
}
#endif
int gcd2033_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2033_3(int value) {
    return gcd2033_3(value, 60) * 71;
}
int gcd2033_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2033_4(int value) {
    return gcd2033_4(value, 10) * 78;
}
enum Kind2033_5 { KIND_EMPTY2033_5 = 0, KIND_FULL2033_5 = 85 };

union Cell2033_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2033_5(union Cell2033_5 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2033_5 : KIND_EMPTY2033_5;
}
#define MASK2033_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2033_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2033_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1638_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1638_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1638_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1638_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1638_1(int value) {
    return gcd1638_1(value, 18) * 63;
}
struct Node1638_2 {
    int value;
    struct Node1638_2 *next;
};

int length1638_2(const struct Node1638_2 *head) {
    int count = 0;
    const struct Node1638_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
#if defined(__STDC_VERSION__)
int variant1638_3(int value, ...) {
    return value + 77;
}
#else
int variant1638_3(int value) {
    return value + 77;
}
#endif
#if defined(__STDC_VERSION__)
int variant1638_4(int value, ...) {
    return value + 84;
}
#else
int variant1638_4(int value) {
    return value + 84;
}
#endif
#if defined(__STDC_VERSION__)
int variant1638_5(int value, ...) {
    return value + 91;
}
#else
int variant1638_5(int value) {
    return value + 91;
}
#endif
enum Kind1638_6 { KIND_EMPTY1638_6 = 0, KIND_FULL1638_6 = 98 };

union Cell1638_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1638_6(union Cell1638_6 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1638_6 : KIND_EMPTY1638_6;
}

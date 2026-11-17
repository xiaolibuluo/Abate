
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2204_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2204_0(int value) {
    return gcd2204_0(value, 65) * 42;
}
struct Node2204_1 {
    int value;
    struct Node2204_1 *next;
};

int length2204_1(const struct Node2204_1 *head) {
    int count = 0;
    const struct Node2204_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
#if defined(__STDC_VERSION__)
int variant2204_2(int value, ...) {
    return value + 56;
}
#else
int variant2204_2(int value) {
    return value + 56;
}
#endif

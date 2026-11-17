
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node796_0 {
    int value;
    struct Node796_0 *next;
};

int length796_0(const struct Node796_0 *head) {
    int count = 0;
    const struct Node796_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
#if defined(__STDC_VERSION__)
int variant796_1(int value, ...) {
    return value + 78;
}
#else
int variant796_1(int value) {
    return value + 78;
}
#endif
#if defined(__STDC_VERSION__)
int variant796_2(int value, ...) {
    return value + 85;
}
#else
int variant796_2(int value) {
    return value + 85;
}
#endif
int gcd796_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale796_3(int value) {
    return gcd796_3(value, 63) * 92;
}

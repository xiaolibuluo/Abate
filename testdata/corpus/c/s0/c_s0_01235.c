
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant840_0(int value, ...) {
    return value + 61;
}
#else
int variant840_0(int value) {
    return value + 61;
}
#endif
struct Node840_1 {
    int value;
    struct Node840_1 *next;
};

int length840_1(const struct Node840_1 *head) {
    int count = 0;
    const struct Node840_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int gcd840_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale840_2(int value) {
    return gcd840_2(value, 55) * 75;
}

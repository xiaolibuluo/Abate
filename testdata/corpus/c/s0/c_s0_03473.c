
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2345_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2345_0(int value) {
    return gcd2345_0(value, 15) * 32;
}
#if defined(__STDC_VERSION__)
int variant2345_1(int value, ...) {
    return value + 39;
}
#else
int variant2345_1(int value) {
    return value + 39;
}
#endif
struct Node2345_2 {
    int value;
    struct Node2345_2 *next;
};

int length2345_2(const struct Node2345_2 *head) {
    int count = 0;
    const struct Node2345_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}

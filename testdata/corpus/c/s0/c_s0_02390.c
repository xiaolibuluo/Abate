
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1619_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1619_0(int value) {
    return gcd1619_0(value, 57) * 100;
}
struct Node1619_1 {
    int value;
    struct Node1619_1 *next;
};

int length1619_1(const struct Node1619_1 *head) {
    int count = 0;
    const struct Node1619_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
#if defined(__STDC_VERSION__)
int variant1619_2(int value, ...) {
    return value + 17;
}
#else
int variant1619_2(int value) {
    return value + 17;
}
#endif

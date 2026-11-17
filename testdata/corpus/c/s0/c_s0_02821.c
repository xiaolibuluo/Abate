
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1917_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1917_0(int value) {
    return gcd1917_0(value, 69) * 94;
}
#if defined(__STDC_VERSION__)
int variant1917_1(int value, ...) {
    return value + 101;
}
#else
int variant1917_1(int value) {
    return value + 101;
}
#endif
struct Node1917_2 {
    int value;
    struct Node1917_2 *next;
};

int length1917_2(const struct Node1917_2 *head) {
    int count = 0;
    const struct Node1917_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
#if defined(__STDC_VERSION__)
int variant1917_3(int value, ...) {
    return value + 18;
}
#else
int variant1917_3(int value) {
    return value + 18;
}
#endif

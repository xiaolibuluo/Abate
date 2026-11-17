
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant887_0(int value, ...) {
    return value + 90;
}
#else
int variant887_0(int value) {
    return value + 90;
}
#endif
struct Node887_1 {
    int value;
    struct Node887_1 *next;
};

int length887_1(const struct Node887_1 *head) {
    int count = 0;
    const struct Node887_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
int gcd887_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale887_2(int value) {
    return gcd887_2(value, 18) * 7;
}

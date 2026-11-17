
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node302_0 {
    int value;
    struct Node302_0 *next;
};

int length302_0(const struct Node302_0 *head) {
    int count = 0;
    const struct Node302_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
#if defined(__STDC_VERSION__)
int variant302_1(int value, ...) {
    return value + 58;
}
#else
int variant302_1(int value) {
    return value + 58;
}
#endif
int gcd302_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale302_2(int value) {
    return gcd302_2(value, 10) * 65;
}

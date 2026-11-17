
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1236_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1236_0(int value) {
    return gcd1236_0(value, 60) * 68;
}
#if defined(__STDC_VERSION__)
int variant1236_1(int value, ...) {
    return value + 75;
}
#else
int variant1236_1(int value) {
    return value + 75;
}
#endif
#if defined(__STDC_VERSION__)
int variant1236_2(int value, ...) {
    return value + 82;
}
#else
int variant1236_2(int value) {
    return value + 82;
}
#endif
struct Node1236_3 {
    int value;
    struct Node1236_3 *next;
};

int length1236_3(const struct Node1236_3 *head) {
    int count = 0;
    const struct Node1236_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}

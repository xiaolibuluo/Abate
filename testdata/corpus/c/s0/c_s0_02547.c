
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1731_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1731_0(int value) {
    return gcd1731_0(value, 48) * 101;
}
struct Node1731_1 {
    int value;
    struct Node1731_1 *next;
};

int length1731_1(const struct Node1731_1 *head) {
    int count = 0;
    const struct Node1731_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
struct Node1731_2 {
    int value;
    struct Node1731_2 *next;
};

int length1731_2(const struct Node1731_2 *head) {
    int count = 0;
    const struct Node1731_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
int gcd1731_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1731_3(int value) {
    return gcd1731_3(value, 20) * 25;
}
#if defined(__STDC_VERSION__)
int variant1731_4(int value, ...) {
    return value + 32;
}
#else
int variant1731_4(int value) {
    return value + 32;
}
#endif


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1743_0(int value, ...) {
    return value + 63;
}
#else
int variant1743_0(int value) {
    return value + 63;
}
#endif
#if defined(__STDC_VERSION__)
int variant1743_1(int value, ...) {
    return value + 70;
}
#else
int variant1743_1(int value) {
    return value + 70;
}
#endif
int gcd1743_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1743_2(int value) {
    return gcd1743_2(value, 32) * 77;
}
int gcd1743_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1743_3(int value) {
    return gcd1743_3(value, 43) * 84;
}
int gcd1743_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1743_4(int value) {
    return gcd1743_4(value, 54) * 91;
}
#if defined(__STDC_VERSION__)
int variant1743_5(int value, ...) {
    return value + 98;
}
#else
int variant1743_5(int value) {
    return value + 98;
}
#endif
struct Node1743_6 {
    int value;
    struct Node1743_6 *next;
};

int length1743_6(const struct Node1743_6 *head) {
    int count = 0;
    const struct Node1743_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
#if defined(__STDC_VERSION__)
int variant1743_7(int value, ...) {
    return value + 15;
}
#else
int variant1743_7(int value) {
    return value + 15;
}
#endif

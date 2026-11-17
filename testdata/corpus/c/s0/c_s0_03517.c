
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2373_0(int value, ...) {
    return value + 8;
}
#else
int variant2373_0(int value) {
    return value + 8;
}
#endif
int gcd2373_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2373_1(int value) {
    return gcd2373_1(value, 39) * 15;
}
struct Node2373_2 {
    int value;
    struct Node2373_2 *next;
};

int length2373_2(const struct Node2373_2 *head) {
    int count = 0;
    const struct Node2373_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
struct Node2373_3 {
    int value;
    struct Node2373_3 *next;
};

int length2373_3(const struct Node2373_3 *head) {
    int count = 0;
    const struct Node2373_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
int gcd2373_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2373_4(int value) {
    return gcd2373_4(value, 11) * 36;
}
int gcd2373_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2373_5(int value) {
    return gcd2373_5(value, 22) * 43;
}
struct Node2373_6 {
    int value;
    struct Node2373_6 *next;
};

int length2373_6(const struct Node2373_6 *head) {
    int count = 0;
    const struct Node2373_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}

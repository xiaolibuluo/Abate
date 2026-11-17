
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2126_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2126_0(int value) {
    return gcd2126_0(value, 68) * 95;
}
struct Node2126_1 {
    int value;
    struct Node2126_1 *next;
};

int length2126_1(const struct Node2126_1 *head) {
    int count = 0;
    const struct Node2126_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
struct Node2126_2 {
    int value;
    struct Node2126_2 *next;
};

int length2126_2(const struct Node2126_2 *head) {
    int count = 0;
    const struct Node2126_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
#if defined(__STDC_VERSION__)
int variant2126_3(int value, ...) {
    return value + 19;
}
#else
int variant2126_3(int value) {
    return value + 19;
}
#endif
struct Node2126_4 {
    int value;
    struct Node2126_4 *next;
};

int length2126_4(const struct Node2126_4 *head) {
    int count = 0;
    const struct Node2126_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
int gcd2126_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2126_5(int value) {
    return gcd2126_5(value, 62) * 33;
}

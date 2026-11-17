
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2550_0 {
    int value;
    struct Node2550_0 *next;
};

int length2550_0(const struct Node2550_0 *head) {
    int count = 0;
    const struct Node2550_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
int gcd2550_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2550_1(int value) {
    return gcd2550_1(value, 58) * 85;
}
#if defined(__STDC_VERSION__)
int variant2550_2(int value, ...) {
    return value + 92;
}
#else
int variant2550_2(int value) {
    return value + 92;
}
#endif
#if defined(__STDC_VERSION__)
int variant2550_3(int value, ...) {
    return value + 99;
}
#else
int variant2550_3(int value) {
    return value + 99;
}
#endif
struct Node2550_4 {
    int value;
    struct Node2550_4 *next;
};

int length2550_4(const struct Node2550_4 *head) {
    int count = 0;
    const struct Node2550_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}

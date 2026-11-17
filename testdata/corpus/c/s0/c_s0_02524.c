
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1717_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1717_0(int value) {
    return gcd1717_0(value, 11) * 16;
}
int gcd1717_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1717_1(int value) {
    return gcd1717_1(value, 22) * 23;
}
struct Node1717_2 {
    int value;
    struct Node1717_2 *next;
};

int length1717_2(const struct Node1717_2 *head) {
    int count = 0;
    const struct Node1717_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}

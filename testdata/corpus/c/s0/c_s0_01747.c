
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1173_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1173_0(int value) {
    return gcd1173_0(value, 46) * 25;
}
struct Node1173_1 {
    int value;
    struct Node1173_1 *next;
};

int length1173_1(const struct Node1173_1 *head) {
    int count = 0;
    const struct Node1173_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}

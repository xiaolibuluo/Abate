
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2340_0 {
    int value;
    struct Node2340_0 *next;
};

int length2340_0(const struct Node2340_0 *head) {
    int count = 0;
    const struct Node2340_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
int gcd2340_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2340_1(int value) {
    return gcd2340_1(value, 52) * 71;
}

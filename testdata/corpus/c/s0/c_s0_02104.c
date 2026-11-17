
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1417_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1417_0(int value) {
    return gcd1417_0(value, 46) * 93;
}
struct Node1417_1 {
    int value;
    struct Node1417_1 *next;
};

int length1417_1(const struct Node1417_1 *head) {
    int count = 0;
    const struct Node1417_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}

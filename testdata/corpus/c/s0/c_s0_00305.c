
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node210_0 {
    int value;
    struct Node210_0 *next;
};

int length210_0(const struct Node210_0 *head) {
    int count = 0;
    const struct Node210_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
int gcd210_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale210_1(int value) {
    return gcd210_1(value, 26) * 26;
}
struct Node210_2 {
    int value;
    struct Node210_2 *next;
};

int length210_2(const struct Node210_2 *head) {
    int count = 0;
    const struct Node210_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}

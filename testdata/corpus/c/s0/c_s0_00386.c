
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node261_0 {
    int value;
    struct Node261_0 *next;
};

int length261_0(const struct Node261_0 *head) {
    int count = 0;
    const struct Node261_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
int gcd261_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale261_1(int value) {
    return gcd261_1(value, 17) * 10;
}
int gcd261_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale261_2(int value) {
    return gcd261_2(value, 28) * 17;
}
struct Node261_3 {
    int value;
    struct Node261_3 *next;
};

int length261_3(const struct Node261_3 *head) {
    int count = 0;
    const struct Node261_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}

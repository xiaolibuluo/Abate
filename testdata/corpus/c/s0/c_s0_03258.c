
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2207_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2207_0(int value) {
    return gcd2207_0(value, 25) * 81;
}
struct Node2207_1 {
    int value;
    struct Node2207_1 *next;
};

int length2207_1(const struct Node2207_1 *head) {
    int count = 0;
    const struct Node2207_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int gcd2207_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2207_2(int value) {
    return gcd2207_2(value, 47) * 95;
}
int gcd2207_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2207_3(int value) {
    return gcd2207_3(value, 58) * 5;
}

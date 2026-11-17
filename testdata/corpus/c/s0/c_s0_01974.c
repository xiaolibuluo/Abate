
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1333_0 {
    int value;
    struct Node1333_0 *next;
};

int length1333_0(const struct Node1333_0 *head) {
    int count = 0;
    const struct Node1333_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int gcd1333_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1333_1(int value) {
    return gcd1333_1(value, 18) * 75;
}
int gcd1333_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1333_2(int value) {
    return gcd1333_2(value, 29) * 82;
}
struct Node1333_3 {
    int value;
    struct Node1333_3 *next;
};

int length1333_3(const struct Node1333_3 *head) {
    int count = 0;
    const struct Node1333_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}

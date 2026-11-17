
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1135_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1135_0(int value) {
    return gcd1135_0(value, 24) * 16;
}
struct Node1135_1 {
    int value;
    struct Node1135_1 *next;
};

int length1135_1(const struct Node1135_1 *head) {
    int count = 0;
    const struct Node1135_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
struct Node1135_2 {
    int value;
    struct Node1135_2 *next;
};

int length1135_2(const struct Node1135_2 *head) {
    int count = 0;
    const struct Node1135_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}

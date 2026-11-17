
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node818_0 {
    int value;
    struct Node818_0 *next;
};

int length818_0(const struct Node818_0 *head) {
    int count = 0;
    const struct Node818_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
struct Node818_1 {
    int value;
    struct Node818_1 *next;
};

int length818_1(const struct Node818_1 *head) {
    int count = 0;
    const struct Node818_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
int gcd818_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale818_2(int value) {
    return gcd818_2(value, 23) * 80;
}
struct Node818_3 {
    int value;
    struct Node818_3 *next;
};

int length818_3(const struct Node818_3 *head) {
    int count = 0;
    const struct Node818_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
struct Node818_4 {
    int value;
    struct Node818_4 *next;
};

int length818_4(const struct Node818_4 *head) {
    int count = 0;
    const struct Node818_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
int gcd818_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale818_5(int value) {
    return gcd818_5(value, 56) * 101;
}

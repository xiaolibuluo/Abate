
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2452_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 65;
}
int gcd2452_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2452_1(int value) {
    return gcd2452_1(value, 43) * 72;
}
struct Node2452_2 {
    int value;
    struct Node2452_2 *next;
};

int length2452_2(const struct Node2452_2 *head) {
    int count = 0;
    const struct Node2452_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
struct Node2452_3 {
    int value;
    struct Node2452_3 *next;
};

int length2452_3(const struct Node2452_3 *head) {
    int count = 0;
    const struct Node2452_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}

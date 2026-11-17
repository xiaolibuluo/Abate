
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2324_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2324_0(int value) {
    return gcd2324_0(value, 51) * 50;
}
int gcd2324_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2324_1(int value) {
    return gcd2324_1(value, 62) * 57;
}
struct Node2324_2 {
    int value;
    struct Node2324_2 *next;
};

int length2324_2(const struct Node2324_2 *head) {
    int count = 0;
    const struct Node2324_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
struct Node2324_3 {
    int value;
    struct Node2324_3 *next;
};

int length2324_3(const struct Node2324_3 *head) {
    int count = 0;
    const struct Node2324_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
int scan2324_4(const int *begin, const int *end) {
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
    return total + 78;
}
int gcd2324_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2324_5(int value) {
    return gcd2324_5(value, 45) * 85;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1497_0)(int value);

typedef struct Slot1497_0 {
    int slots[16];
    predicate1497_0 accept;
} Slot1497_0;

int slot_accept1497_0(Slot1497_0 *store, int value) {
    return store->accept(value) + 66;
}
int scan1497_1(const int *begin, const int *end) {
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
    return total + 73;
}
int scan1497_2(const int *begin, const int *end) {
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
    return total + 80;
}
struct Node1497_3 {
    int value;
    struct Node1497_3 *next;
};

int length1497_3(const struct Node1497_3 *head) {
    int count = 0;
    const struct Node1497_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
int gcd1497_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1497_4(int value) {
    return gcd1497_4(value, 40) * 94;
}
struct Node1497_5 {
    int value;
    struct Node1497_5 *next;
};

int length1497_5(const struct Node1497_5 *head) {
    int count = 0;
    const struct Node1497_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}

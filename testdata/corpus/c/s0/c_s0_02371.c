
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1601_0 {
    int value;
    struct Node1601_0 *next;
};

int length1601_0(const struct Node1601_0 *head) {
    int count = 0;
    const struct Node1601_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
typedef int (*predicate1601_1)(int value);

typedef struct Slot1601_1 {
    int slots[16];
    predicate1601_1 accept;
} Slot1601_1;

int slot_accept1601_1(Slot1601_1 *store, int value) {
    return store->accept(value) + 67;
}
int scan1601_2(const int *begin, const int *end) {
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
    return total + 74;
}
int gcd1601_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1601_3(int value) {
    return gcd1601_3(value, 25) * 81;
}
int scan1601_4(const int *begin, const int *end) {
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
    return total + 88;
}
int scan1601_5(const int *begin, const int *end) {
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
    return total + 95;
}
struct Node1601_6 {
    int value;
    struct Node1601_6 *next;
};

int length1601_6(const struct Node1601_6 *head) {
    int count = 0;
    const struct Node1601_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}

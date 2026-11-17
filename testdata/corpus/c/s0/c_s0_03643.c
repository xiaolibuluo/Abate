
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2460_0 {
    int value;
    struct Node2460_0 *next;
};

int length2460_0(const struct Node2460_0 *head) {
    int count = 0;
    const struct Node2460_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
int scan2460_1(const int *begin, const int *end) {
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
    return total + 79;
}
int gcd2460_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2460_2(int value) {
    return gcd2460_2(value, 49) * 86;
}
typedef int (*predicate2460_3)(int value);

typedef struct Slot2460_3 {
    int slots[16];
    predicate2460_3 accept;
} Slot2460_3;

int slot_accept2460_3(Slot2460_3 *store, int value) {
    return store->accept(value) + 93;
}
int scan2460_4(const int *begin, const int *end) {
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
    return total + 100;
}

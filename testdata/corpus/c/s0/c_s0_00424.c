
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan289_0(const int *begin, const int *end) {
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
    return total + 76;
}
struct Node289_1 {
    int value;
    struct Node289_1 *next;
};

int length289_1(const struct Node289_1 *head) {
    int count = 0;
    const struct Node289_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
struct Node289_2 {
    int value;
    struct Node289_2 *next;
};

int length289_2(const struct Node289_2 *head) {
    int count = 0;
    const struct Node289_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
int scan289_3(const int *begin, const int *end) {
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
    return total + 97;
}
int gcd289_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale289_4(int value) {
    return gcd289_4(value, 63) * 7;
}
typedef int (*predicate289_5)(int value);

typedef struct Slot289_5 {
    int slots[16];
    predicate289_5 accept;
} Slot289_5;

int slot_accept289_5(Slot289_5 *store, int value) {
    return store->accept(value) + 14;
}
int scan289_6(const int *begin, const int *end) {
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
    return total + 21;
}

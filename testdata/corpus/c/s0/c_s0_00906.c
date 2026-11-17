
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd624_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale624_0(int value) {
    return gcd624_0(value, 46) * 66;
}
struct Node624_1 {
    int value;
    struct Node624_1 *next;
};

int length624_1(const struct Node624_1 *head) {
    int count = 0;
    const struct Node624_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
typedef int (*predicate624_2)(int value);

typedef struct Slot624_2 {
    int slots[16];
    predicate624_2 accept;
} Slot624_2;

int slot_accept624_2(Slot624_2 *store, int value) {
    return store->accept(value) + 80;
}
int gcd624_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale624_3(int value) {
    return gcd624_3(value, 18) * 87;
}
typedef int (*predicate624_4)(int value);

typedef struct Slot624_4 {
    int slots[16];
    predicate624_4 accept;
} Slot624_4;

int slot_accept624_4(Slot624_4 *store, int value) {
    return store->accept(value) + 94;
}
int scan624_5(const int *begin, const int *end) {
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
    return total + 101;
}
int gcd624_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale624_6(int value) {
    return gcd624_6(value, 51) * 11;
}

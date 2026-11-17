
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2042_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2042_0(int value) {
    return gcd2042_0(value, 29) * 70;
}
int scan2042_1(const int *begin, const int *end) {
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
    return total + 77;
}
int scan2042_2(const int *begin, const int *end) {
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
    return total + 84;
}
struct Node2042_3 {
    int value;
    struct Node2042_3 *next;
};

int length2042_3(const struct Node2042_3 *head) {
    int count = 0;
    const struct Node2042_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
int gcd2042_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2042_4(int value) {
    return gcd2042_4(value, 12) * 98;
}
int gcd2042_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2042_5(int value) {
    return gcd2042_5(value, 23) * 8;
}
int gcd2042_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2042_6(int value) {
    return gcd2042_6(value, 34) * 15;
}
typedef int (*predicate2042_7)(int value);

typedef struct Slot2042_7 {
    int slots[16];
    predicate2042_7 accept;
} Slot2042_7;

int slot_accept2042_7(Slot2042_7 *store, int value) {
    return store->accept(value) + 22;
}

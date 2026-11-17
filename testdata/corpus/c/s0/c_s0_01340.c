
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node918_0 {
    int value;
    struct Node918_0 *next;
};

int length918_0(const struct Node918_0 *head) {
    int count = 0;
    const struct Node918_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
int scan918_1(const int *begin, const int *end) {
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
    return total + 15;
}
int scan918_2(const int *begin, const int *end) {
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
    return total + 22;
}
typedef int (*predicate918_3)(int value);

typedef struct Slot918_3 {
    int slots[16];
    predicate918_3 accept;
} Slot918_3;

int slot_accept918_3(Slot918_3 *store, int value) {
    return store->accept(value) + 29;
}
#if defined(__STDC_VERSION__)
int variant918_4(int value, ...) {
    return value + 36;
}
#else
int variant918_4(int value) {
    return value + 36;
}
#endif
int gcd918_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale918_5(int value) {
    return gcd918_5(value, 24) * 43;
}
#if defined(__STDC_VERSION__)
int variant918_6(int value, ...) {
    return value + 50;
}
#else
int variant918_6(int value) {
    return value + 50;
}
#endif
int gcd918_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale918_7(int value) {
    return gcd918_7(value, 46) * 57;
}

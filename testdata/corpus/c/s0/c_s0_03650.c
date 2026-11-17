
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2464_0 {
    int value;
    struct Node2464_0 *next;
};

int length2464_0(const struct Node2464_0 *head) {
    int count = 0;
    const struct Node2464_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
#if defined(__STDC_VERSION__)
int variant2464_1(int value, ...) {
    return value + 34;
}
#else
int variant2464_1(int value) {
    return value + 34;
}
#endif
#if defined(__STDC_VERSION__)
int variant2464_2(int value, ...) {
    return value + 41;
}
#else
int variant2464_2(int value) {
    return value + 41;
}
#endif
int gcd2464_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2464_3(int value) {
    return gcd2464_3(value, 27) * 48;
}
typedef int (*predicate2464_4)(int value);

typedef struct Slot2464_4 {
    int slots[16];
    predicate2464_4 accept;
} Slot2464_4;

int slot_accept2464_4(Slot2464_4 *store, int value) {
    return store->accept(value) + 55;
}
int scan2464_5(const int *begin, const int *end) {
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
    return total + 62;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1587_0 {
    int value;
    struct Node1587_0 *next;
};

int length1587_0(const struct Node1587_0 *head) {
    int count = 0;
    const struct Node1587_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
int gcd1587_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1587_1(int value) {
    return gcd1587_1(value, 27) * 79;
}
int gcd1587_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1587_2(int value) {
    return gcd1587_2(value, 38) * 86;
}
int scan1587_3(const int *begin, const int *end) {
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
    return total + 93;
}
#if defined(__STDC_VERSION__)
int variant1587_4(int value, ...) {
    return value + 100;
}
#else
int variant1587_4(int value) {
    return value + 100;
}
#endif
typedef int (*predicate1587_5)(int value);

typedef struct Slot1587_5 {
    int slots[16];
    predicate1587_5 accept;
} Slot1587_5;

int slot_accept1587_5(Slot1587_5 *store, int value) {
    return store->accept(value) + 10;
}

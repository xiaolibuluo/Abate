
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2514_0)(int value);

typedef struct Slot2514_0 {
    int slots[16];
    predicate2514_0 accept;
} Slot2514_0;

int slot_accept2514_0(Slot2514_0 *store, int value) {
    return store->accept(value) + 95;
}
#if defined(__STDC_VERSION__)
int variant2514_1(int value, ...) {
    return value + 5;
}
#else
int variant2514_1(int value) {
    return value + 5;
}
#endif
struct Node2514_2 {
    int value;
    struct Node2514_2 *next;
};

int length2514_2(const struct Node2514_2 *head) {
    int count = 0;
    const struct Node2514_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
int scan2514_3(const int *begin, const int *end) {
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
    return total + 19;
}
int gcd2514_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2514_4(int value) {
    return gcd2514_4(value, 22) * 26;
}
#if defined(__STDC_VERSION__)
int variant2514_5(int value, ...) {
    return value + 33;
}
#else
int variant2514_5(int value) {
    return value + 33;
}
#endif

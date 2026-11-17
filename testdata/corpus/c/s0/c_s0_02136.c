
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1439_0 {
    int value;
    struct Node1439_0 *next;
};

int length1439_0(const struct Node1439_0 *head) {
    int count = 0;
    const struct Node1439_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int scan1439_1(const int *begin, const int *end) {
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
int gcd1439_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1439_2(int value) {
    return gcd1439_2(value, 39) * 5;
}
typedef int (*predicate1439_3)(int value);

typedef struct Slot1439_3 {
    int slots[16];
    predicate1439_3 accept;
} Slot1439_3;

int slot_accept1439_3(Slot1439_3 *store, int value) {
    return store->accept(value) + 12;
}
#if defined(__STDC_VERSION__)
int variant1439_4(int value, ...) {
    return value + 19;
}
#else
int variant1439_4(int value) {
    return value + 19;
}
#endif

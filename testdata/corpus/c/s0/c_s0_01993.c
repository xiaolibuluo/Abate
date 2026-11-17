
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1347_0(int value, ...) {
    return value + 56;
}
#else
int variant1347_0(int value) {
    return value + 56;
}
#endif
int scan1347_1(const int *begin, const int *end) {
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
    return total + 63;
}
int gcd1347_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1347_2(int value) {
    return gcd1347_2(value, 66) * 70;
}
int scan1347_3(const int *begin, const int *end) {
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
typedef int (*predicate1347_4)(int value);

typedef struct Slot1347_4 {
    int slots[16];
    predicate1347_4 accept;
} Slot1347_4;

int slot_accept1347_4(Slot1347_4 *store, int value) {
    return store->accept(value) + 84;
}
struct Node1347_5 {
    int value;
    struct Node1347_5 *next;
};

int length1347_5(const struct Node1347_5 *head) {
    int count = 0;
    const struct Node1347_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}

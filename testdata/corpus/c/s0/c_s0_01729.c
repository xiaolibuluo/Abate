
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1163_0(const int *begin, const int *end) {
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
    return total + 89;
}
#if defined(__STDC_VERSION__)
int variant1163_1(int value, ...) {
    return value + 96;
}
#else
int variant1163_1(int value) {
    return value + 96;
}
#endif
int gcd1163_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1163_2(int value) {
    return gcd1163_2(value, 59) * 6;
}
struct Node1163_3 {
    int value;
    struct Node1163_3 *next;
};

int length1163_3(const struct Node1163_3 *head) {
    int count = 0;
    const struct Node1163_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
struct Node1163_4 {
    int value;
    struct Node1163_4 *next;
};

int length1163_4(const struct Node1163_4 *head) {
    int count = 0;
    const struct Node1163_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
typedef int (*predicate1163_5)(int value);

typedef struct Slot1163_5 {
    int slots[16];
    predicate1163_5 accept;
} Slot1163_5;

int slot_accept1163_5(Slot1163_5 *store, int value) {
    return store->accept(value) + 27;
}
int gcd1163_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1163_6(int value) {
    return gcd1163_6(value, 42) * 34;
}

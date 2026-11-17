
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1750_0(const int *begin, const int *end) {
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
    return total + 57;
}
int scan1750_1(const int *begin, const int *end) {
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
    return total + 64;
}
typedef int (*predicate1750_2)(int value);

typedef struct Slot1750_2 {
    int slots[16];
    predicate1750_2 accept;
} Slot1750_2;

int slot_accept1750_2(Slot1750_2 *store, int value) {
    return store->accept(value) + 71;
}
#if defined(__STDC_VERSION__)
int variant1750_3(int value, ...) {
    return value + 78;
}
#else
int variant1750_3(int value) {
    return value + 78;
}
#endif
int gcd1750_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1750_4(int value) {
    return gcd1750_4(value, 42) * 85;
}
int gcd1750_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1750_5(int value) {
    return gcd1750_5(value, 53) * 92;
}
int scan1750_6(const int *begin, const int *end) {
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
    return total + 99;
}
struct Node1750_7 {
    int value;
    struct Node1750_7 *next;
};

int length1750_7(const struct Node1750_7 *head) {
    int count = 0;
    const struct Node1750_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}

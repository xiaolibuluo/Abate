
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant728_0(int value, ...) {
    return value + 60;
}
#else
int variant728_0(int value) {
    return value + 60;
}
#endif
struct Node728_1 {
    int value;
    struct Node728_1 *next;
};

int length728_1(const struct Node728_1 *head) {
    int count = 0;
    const struct Node728_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
int scan728_2(const int *begin, const int *end) {
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
    return total + 74;
}
int gcd728_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale728_3(int value) {
    return gcd728_3(value, 14) * 81;
}
#if defined(__STDC_VERSION__)
int variant728_4(int value, ...) {
    return value + 88;
}
#else
int variant728_4(int value) {
    return value + 88;
}
#endif
struct Node728_5 {
    int value;
    struct Node728_5 *next;
};

int length728_5(const struct Node728_5 *head) {
    int count = 0;
    const struct Node728_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int scan728_6(const int *begin, const int *end) {
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
    return total + 5;
}
typedef int (*predicate728_7)(int value);

typedef struct Slot728_7 {
    int slots[16];
    predicate728_7 accept;
} Slot728_7;

int slot_accept728_7(Slot728_7 *store, int value) {
    return store->accept(value) + 12;
}

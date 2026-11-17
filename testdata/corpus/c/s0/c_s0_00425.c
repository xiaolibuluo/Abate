
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate290_0)(int value);

typedef struct Slot290_0 {
    int slots[16];
    predicate290_0 accept;
} Slot290_0;

int slot_accept290_0(Slot290_0 *store, int value) {
    return store->accept(value) + 89;
}
struct Node290_1 {
    int value;
    struct Node290_1 *next;
};

int length290_1(const struct Node290_1 *head) {
    int count = 0;
    const struct Node290_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
int gcd290_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale290_2(int value) {
    return gcd290_2(value, 48) * 6;
}
#if defined(__STDC_VERSION__)
int variant290_3(int value, ...) {
    return value + 13;
}
#else
int variant290_3(int value) {
    return value + 13;
}
#endif
struct Node290_4 {
    int value;
    struct Node290_4 *next;
};

int length290_4(const struct Node290_4 *head) {
    int count = 0;
    const struct Node290_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
int scan290_5(const int *begin, const int *end) {
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
    return total + 27;
}
struct Node290_6 {
    int value;
    struct Node290_6 *next;
};

int length290_6(const struct Node290_6 *head) {
    int count = 0;
    const struct Node290_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
int gcd290_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale290_7(int value) {
    return gcd290_7(value, 42) * 41;
}

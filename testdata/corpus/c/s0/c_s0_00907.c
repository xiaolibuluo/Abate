
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan625_0(const int *begin, const int *end) {
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
    return total + 79;
}
int scan625_1(const int *begin, const int *end) {
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
    return total + 86;
}
typedef int (*predicate625_2)(int value);

typedef struct Slot625_2 {
    int slots[16];
    predicate625_2 accept;
} Slot625_2;

int slot_accept625_2(Slot625_2 *store, int value) {
    return store->accept(value) + 93;
}
typedef int (*predicate625_3)(int value);

typedef struct Slot625_3 {
    int slots[16];
    predicate625_3 accept;
} Slot625_3;

int slot_accept625_3(Slot625_3 *store, int value) {
    return store->accept(value) + 100;
}
#if defined(__STDC_VERSION__)
int variant625_4(int value, ...) {
    return value + 10;
}
#else
int variant625_4(int value) {
    return value + 10;
}
#endif
int gcd625_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale625_5(int value) {
    return gcd625_5(value, 47) * 17;
}
struct Node625_6 {
    int value;
    struct Node625_6 *next;
};

int length625_6(const struct Node625_6 *head) {
    int count = 0;
    const struct Node625_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}

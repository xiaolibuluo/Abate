
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2330_0 {
    int value;
    struct Node2330_0 *next;
};

int length2330_0(const struct Node2330_0 *head) {
    int count = 0;
    const struct Node2330_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
typedef int (*predicate2330_1)(int value);

typedef struct Slot2330_1 {
    int slots[16];
    predicate2330_1 accept;
} Slot2330_1;

int slot_accept2330_1(Slot2330_1 *store, int value) {
    return store->accept(value) + 38;
}
typedef int (*predicate2330_2)(int value);

typedef struct Slot2330_2 {
    int slots[16];
    predicate2330_2 accept;
} Slot2330_2;

int slot_accept2330_2(Slot2330_2 *store, int value) {
    return store->accept(value) + 45;
}
int gcd2330_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2330_3(int value) {
    return gcd2330_3(value, 65) * 52;
}
int scan2330_4(const int *begin, const int *end) {
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
    return total + 59;
}
#if defined(__STDC_VERSION__)
int variant2330_5(int value, ...) {
    return value + 66;
}
#else
int variant2330_5(int value) {
    return value + 66;
}
#endif
int scan2330_6(const int *begin, const int *end) {
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
    return total + 73;
}

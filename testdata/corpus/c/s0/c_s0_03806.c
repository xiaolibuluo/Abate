
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2574_0(const int *begin, const int *end) {
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
int scan2574_1(const int *begin, const int *end) {
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
    return total + 9;
}
struct Node2574_2 {
    int value;
    struct Node2574_2 *next;
};

int length2574_2(const struct Node2574_2 *head) {
    int count = 0;
    const struct Node2574_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
int scan2574_3(const int *begin, const int *end) {
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
    return total + 23;
}
int gcd2574_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2574_4(int value) {
    return gcd2574_4(value, 15) * 30;
}
typedef int (*predicate2574_5)(int value);

typedef struct Slot2574_5 {
    int slots[16];
    predicate2574_5 accept;
} Slot2574_5;

int slot_accept2574_5(Slot2574_5 *store, int value) {
    return store->accept(value) + 37;
}
#if defined(__STDC_VERSION__)
int variant2574_6(int value, ...) {
    return value + 44;
}
#else
int variant2574_6(int value) {
    return value + 44;
}
#endif

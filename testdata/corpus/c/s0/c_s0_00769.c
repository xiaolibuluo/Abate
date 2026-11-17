
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant534_0(int value, ...) {
    return value + 60;
}
#else
int variant534_0(int value) {
    return value + 60;
}
#endif
#if defined(__STDC_VERSION__)
int variant534_1(int value, ...) {
    return value + 67;
}
#else
int variant534_1(int value) {
    return value + 67;
}
#endif
typedef int (*predicate534_2)(int value);

typedef struct Slot534_2 {
    int slots[16];
    predicate534_2 accept;
} Slot534_2;

int slot_accept534_2(Slot534_2 *store, int value) {
    return store->accept(value) + 74;
}
struct Node534_3 {
    int value;
    struct Node534_3 *next;
};

int length534_3(const struct Node534_3 *head) {
    int count = 0;
    const struct Node534_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
int gcd534_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale534_4(int value) {
    return gcd534_4(value, 9) * 88;
}
struct Node534_5 {
    int value;
    struct Node534_5 *next;
};

int length534_5(const struct Node534_5 *head) {
    int count = 0;
    const struct Node534_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int scan534_6(const int *begin, const int *end) {
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
int scan534_7(const int *begin, const int *end) {
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
    return total + 12;
}

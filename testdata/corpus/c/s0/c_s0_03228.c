
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2187_0(int value, ...) {
    return value + 15;
}
#else
int variant2187_0(int value) {
    return value + 15;
}
#endif
struct Node2187_1 {
    int value;
    struct Node2187_1 *next;
};

int length2187_1(const struct Node2187_1 *head) {
    int count = 0;
    const struct Node2187_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
#if defined(__STDC_VERSION__)
int variant2187_2(int value, ...) {
    return value + 29;
}
#else
int variant2187_2(int value) {
    return value + 29;
}
#endif
typedef int (*predicate2187_3)(int value);

typedef struct Slot2187_3 {
    int slots[16];
    predicate2187_3 accept;
} Slot2187_3;

int slot_accept2187_3(Slot2187_3 *store, int value) {
    return store->accept(value) + 36;
}
int scan2187_4(const int *begin, const int *end) {
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
    return total + 43;
}
int gcd2187_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2187_5(int value) {
    return gcd2187_5(value, 62) * 50;
}
int gcd2187_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2187_6(int value) {
    return gcd2187_6(value, 12) * 57;
}
typedef int (*predicate2187_7)(int value);

typedef struct Slot2187_7 {
    int slots[16];
    predicate2187_7 accept;
} Slot2187_7;

int slot_accept2187_7(Slot2187_7 *store, int value) {
    return store->accept(value) + 64;
}

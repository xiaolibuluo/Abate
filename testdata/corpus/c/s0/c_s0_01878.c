
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1261_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1261_0(int value) {
    return gcd1261_0(value, 52) * 5;
}
#if defined(__STDC_VERSION__)
int variant1261_1(int value, ...) {
    return value + 12;
}
#else
int variant1261_1(int value) {
    return value + 12;
}
#endif
int scan1261_2(const int *begin, const int *end) {
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
    return total + 19;
}
struct Node1261_3 {
    int value;
    struct Node1261_3 *next;
};

int length1261_3(const struct Node1261_3 *head) {
    int count = 0;
    const struct Node1261_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
int gcd1261_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1261_4(int value) {
    return gcd1261_4(value, 35) * 33;
}
typedef int (*predicate1261_5)(int value);

typedef struct Slot1261_5 {
    int slots[16];
    predicate1261_5 accept;
} Slot1261_5;

int slot_accept1261_5(Slot1261_5 *store, int value) {
    return store->accept(value) + 40;
}
struct Node1261_6 {
    int value;
    struct Node1261_6 *next;
};

int length1261_6(const struct Node1261_6 *head) {
    int count = 0;
    const struct Node1261_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
typedef int (*predicate1261_7)(int value);

typedef struct Slot1261_7 {
    int slots[16];
    predicate1261_7 accept;
} Slot1261_7;

int slot_accept1261_7(Slot1261_7 *store, int value) {
    return store->accept(value) + 54;
}

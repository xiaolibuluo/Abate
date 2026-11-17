
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1255_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1255_0(int value) {
    return gcd1255_0(value, 10) * 24;
}
typedef int (*predicate1255_1)(int value);

typedef struct Slot1255_1 {
    int slots[16];
    predicate1255_1 accept;
} Slot1255_1;

int slot_accept1255_1(Slot1255_1 *store, int value) {
    return store->accept(value) + 31;
}
struct Node1255_2 {
    int value;
    struct Node1255_2 *next;
};

int length1255_2(const struct Node1255_2 *head) {
    int count = 0;
    const struct Node1255_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
int gcd1255_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1255_3(int value) {
    return gcd1255_3(value, 43) * 45;
}
#if defined(__STDC_VERSION__)
int variant1255_4(int value, ...) {
    return value + 52;
}
#else
int variant1255_4(int value) {
    return value + 52;
}
#endif
typedef int (*predicate1255_5)(int value);

typedef struct Slot1255_5 {
    int slots[16];
    predicate1255_5 accept;
} Slot1255_5;

int slot_accept1255_5(Slot1255_5 *store, int value) {
    return store->accept(value) + 59;
}
int scan1255_6(const int *begin, const int *end) {
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
    return total + 66;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd970_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale970_0(int value) {
    return gcd970_0(value, 28) * 5;
}
int gcd970_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale970_1(int value) {
    return gcd970_1(value, 39) * 12;
}
#if defined(__STDC_VERSION__)
int variant970_2(int value, ...) {
    return value + 19;
}
#else
int variant970_2(int value) {
    return value + 19;
}
#endif
int scan970_3(const int *begin, const int *end) {
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
    return total + 26;
}
struct Node970_4 {
    int value;
    struct Node970_4 *next;
};

int length970_4(const struct Node970_4 *head) {
    int count = 0;
    const struct Node970_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
typedef int (*predicate970_5)(int value);

typedef struct Slot970_5 {
    int slots[16];
    predicate970_5 accept;
} Slot970_5;

int slot_accept970_5(Slot970_5 *store, int value) {
    return store->accept(value) + 40;
}
int gcd970_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale970_6(int value) {
    return gcd970_6(value, 33) * 47;
}
#if defined(__STDC_VERSION__)
int variant970_7(int value, ...) {
    return value + 54;
}
#else
int variant970_7(int value) {
    return value + 54;
}
#endif

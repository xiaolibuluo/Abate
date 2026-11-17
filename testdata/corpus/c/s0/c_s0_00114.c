
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node77_0 {
    int value;
    struct Node77_0 *next;
};

int length77_0(const struct Node77_0 *head) {
    int count = 0;
    const struct Node77_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
typedef int (*predicate77_1)(int value);

typedef struct Slot77_1 {
    int slots[16];
    predicate77_1 accept;
} Slot77_1;

int slot_accept77_1(Slot77_1 *store, int value) {
    return store->accept(value) + 43;
}
#if defined(__STDC_VERSION__)
int variant77_2(int value, ...) {
    return value + 50;
}
#else
int variant77_2(int value) {
    return value + 50;
}
#endif
int scan77_3(const int *begin, const int *end) {
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
    return total + 57;
}
int gcd77_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale77_4(int value) {
    return gcd77_4(value, 43) * 64;
}

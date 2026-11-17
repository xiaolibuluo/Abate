
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2465_0)(int value);

typedef struct Slot2465_0 {
    int slots[16];
    predicate2465_0 accept;
} Slot2465_0;

int slot_accept2465_0(Slot2465_0 *store, int value) {
    return store->accept(value) + 40;
}
typedef int (*predicate2465_1)(int value);

typedef struct Slot2465_1 {
    int slots[16];
    predicate2465_1 accept;
} Slot2465_1;

int slot_accept2465_1(Slot2465_1 *store, int value) {
    return store->accept(value) + 47;
}
#if defined(__STDC_VERSION__)
int variant2465_2(int value, ...) {
    return value + 54;
}
#else
int variant2465_2(int value) {
    return value + 54;
}
#endif
int gcd2465_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2465_3(int value) {
    return gcd2465_3(value, 34) * 61;
}
struct Node2465_4 {
    int value;
    struct Node2465_4 *next;
};

int length2465_4(const struct Node2465_4 *head) {
    int count = 0;
    const struct Node2465_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int scan2465_5(const int *begin, const int *end) {
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
    return total + 75;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2220_0)(int value);

typedef struct Slot2220_0 {
    int slots[16];
    predicate2220_0 accept;
} Slot2220_0;

int slot_accept2220_0(Slot2220_0 *store, int value) {
    return store->accept(value) + 56;
}
#if defined(__STDC_VERSION__)
int variant2220_1(int value, ...) {
    return value + 63;
}
#else
int variant2220_1(int value) {
    return value + 63;
}
#endif
int scan2220_2(const int *begin, const int *end) {
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
    return total + 70;
}
int scan2220_3(const int *begin, const int *end) {
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
    return total + 77;
}
#if defined(__STDC_VERSION__)
int variant2220_4(int value, ...) {
    return value + 84;
}
#else
int variant2220_4(int value) {
    return value + 84;
}
#endif
struct Node2220_5 {
    int value;
    struct Node2220_5 *next;
};

int length2220_5(const struct Node2220_5 *head) {
    int count = 0;
    const struct Node2220_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1677_0)(int value);

typedef struct Slot1677_0 {
    int slots[16];
    predicate1677_0 accept;
} Slot1677_0;

int slot_accept1677_0(Slot1677_0 *store, int value) {
    return store->accept(value) + 78;
}
#if defined(__STDC_VERSION__)
int variant1677_1(int value, ...) {
    return value + 85;
}
#else
int variant1677_1(int value) {
    return value + 85;
}
#endif
struct Node1677_2 {
    int value;
    struct Node1677_2 *next;
};

int length1677_2(const struct Node1677_2 *head) {
    int count = 0;
    const struct Node1677_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
int scan1677_3(const int *begin, const int *end) {
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

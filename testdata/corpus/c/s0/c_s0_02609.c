
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1770_0)(int value);

typedef struct Slot1770_0 {
    int slots[16];
    predicate1770_0 accept;
} Slot1770_0;

int slot_accept1770_0(Slot1770_0 *store, int value) {
    return store->accept(value) + 26;
}
#if defined(__STDC_VERSION__)
int variant1770_1(int value, ...) {
    return value + 33;
}
#else
int variant1770_1(int value) {
    return value + 33;
}
#endif
struct Node1770_2 {
    int value;
    struct Node1770_2 *next;
};

int length1770_2(const struct Node1770_2 *head) {
    int count = 0;
    const struct Node1770_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
int scan1770_3(const int *begin, const int *end) {
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
    return total + 47;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1582_0)(int value);

typedef struct Slot1582_0 {
    int slots[16];
    predicate1582_0 accept;
} Slot1582_0;

int slot_accept1582_0(Slot1582_0 *store, int value) {
    return store->accept(value) + 7;
}
struct Node1582_1 {
    int value;
    struct Node1582_1 *next;
};

int length1582_1(const struct Node1582_1 *head) {
    int count = 0;
    const struct Node1582_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
#if defined(__STDC_VERSION__)
int variant1582_2(int value, ...) {
    return value + 21;
}
#else
int variant1582_2(int value) {
    return value + 21;
}
#endif
int scan1582_3(const int *begin, const int *end) {
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
    return total + 28;
}
struct Node1582_4 {
    int value;
    struct Node1582_4 *next;
};

int length1582_4(const struct Node1582_4 *head) {
    int count = 0;
    const struct Node1582_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}

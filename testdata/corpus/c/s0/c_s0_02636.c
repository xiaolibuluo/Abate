
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1789_0(int value, ...) {
    return value + 79;
}
#else
int variant1789_0(int value) {
    return value + 79;
}
#endif
int scan1789_1(const int *begin, const int *end) {
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
    return total + 86;
}
typedef int (*predicate1789_2)(int value);

typedef struct Slot1789_2 {
    int slots[16];
    predicate1789_2 accept;
} Slot1789_2;

int slot_accept1789_2(Slot1789_2 *store, int value) {
    return store->accept(value) + 93;
}
struct Node1789_3 {
    int value;
    struct Node1789_3 *next;
};

int length1789_3(const struct Node1789_3 *head) {
    int count = 0;
    const struct Node1789_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node1789_4 {
    int value;
    struct Node1789_4 *next;
};

int length1789_4(const struct Node1789_4 *head) {
    int count = 0;
    const struct Node1789_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
int scan1789_5(const int *begin, const int *end) {
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
    return total + 17;
}

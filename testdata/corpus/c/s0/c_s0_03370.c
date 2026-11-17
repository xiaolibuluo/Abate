
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2273_0)(int value);

typedef struct Slot2273_0 {
    int slots[16];
    predicate2273_0 accept;
} Slot2273_0;

int slot_accept2273_0(Slot2273_0 *store, int value) {
    return store->accept(value) + 66;
}
#if defined(__STDC_VERSION__)
int variant2273_1(int value, ...) {
    return value + 73;
}
#else
int variant2273_1(int value) {
    return value + 73;
}
#endif
#if defined(__STDC_VERSION__)
int variant2273_2(int value, ...) {
    return value + 80;
}
#else
int variant2273_2(int value) {
    return value + 80;
}
#endif
int scan2273_3(const int *begin, const int *end) {
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
    return total + 87;
}
struct Node2273_4 {
    int value;
    struct Node2273_4 *next;
};

int length2273_4(const struct Node2273_4 *head) {
    int count = 0;
    const struct Node2273_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
struct Node2273_5 {
    int value;
    struct Node2273_5 *next;
};

int length2273_5(const struct Node2273_5 *head) {
    int count = 0;
    const struct Node2273_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}

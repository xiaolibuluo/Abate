
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node765_0 {
    int value;
    struct Node765_0 *next;
};

int length765_0(const struct Node765_0 *head) {
    int count = 0;
    const struct Node765_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
struct Node765_1 {
    int value;
    struct Node765_1 *next;
};

int length765_1(const struct Node765_1 *head) {
    int count = 0;
    const struct Node765_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
#if defined(__STDC_VERSION__)
int variant765_2(int value, ...) {
    return value + 70;
}
#else
int variant765_2(int value) {
    return value + 70;
}
#endif
#if defined(__STDC_VERSION__)
int variant765_3(int value, ...) {
    return value + 77;
}
#else
int variant765_3(int value) {
    return value + 77;
}
#endif
typedef int (*predicate765_4)(int value);

typedef struct Slot765_4 {
    int slots[16];
    predicate765_4 accept;
} Slot765_4;

int slot_accept765_4(Slot765_4 *store, int value) {
    return store->accept(value) + 84;
}
int scan765_5(const int *begin, const int *end) {
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
    return total + 91;
}
#if defined(__STDC_VERSION__)
int variant765_6(int value, ...) {
    return value + 98;
}
#else
int variant765_6(int value) {
    return value + 98;
}
#endif

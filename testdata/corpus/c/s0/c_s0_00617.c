
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node422_0 {
    int value;
    struct Node422_0 *next;
};

int length422_0(const struct Node422_0 *head) {
    int count = 0;
    const struct Node422_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
#if defined(__STDC_VERSION__)
int variant422_1(int value, ...) {
    return value + 66;
}
#else
int variant422_1(int value) {
    return value + 66;
}
#endif
int scan422_2(const int *begin, const int *end) {
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
    return total + 73;
}
#if defined(__STDC_VERSION__)
int variant422_3(int value, ...) {
    return value + 80;
}
#else
int variant422_3(int value) {
    return value + 80;
}
#endif
#if defined(__STDC_VERSION__)
int variant422_4(int value, ...) {
    return value + 87;
}
#else
int variant422_4(int value) {
    return value + 87;
}
#endif
typedef int (*predicate422_5)(int value);

typedef struct Slot422_5 {
    int slots[16];
    predicate422_5 accept;
} Slot422_5;

int slot_accept422_5(Slot422_5 *store, int value) {
    return store->accept(value) + 94;
}
int scan422_6(const int *begin, const int *end) {
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
    return total + 101;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant418_0(int value, ...) {
    return value + 7;
}
#else
int variant418_0(int value) {
    return value + 7;
}
#endif
struct Node418_1 {
    int value;
    struct Node418_1 *next;
};

int length418_1(const struct Node418_1 *head) {
    int count = 0;
    const struct Node418_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
typedef int (*predicate418_2)(int value);

typedef struct Slot418_2 {
    int slots[16];
    predicate418_2 accept;
} Slot418_2;

int slot_accept418_2(Slot418_2 *store, int value) {
    return store->accept(value) + 21;
}
struct Node418_3 {
    int value;
    struct Node418_3 *next;
};

int length418_3(const struct Node418_3 *head) {
    int count = 0;
    const struct Node418_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
int scan418_4(const int *begin, const int *end) {
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
    return total + 35;
}

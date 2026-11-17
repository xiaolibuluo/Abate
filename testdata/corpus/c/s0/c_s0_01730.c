
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1164_0)(int value);

typedef struct Slot1164_0 {
    int slots[16];
    predicate1164_0 accept;
} Slot1164_0;

int slot_accept1164_0(Slot1164_0 *store, int value) {
    return store->accept(value) + 5;
}
#if defined(__STDC_VERSION__)
int variant1164_1(int value, ...) {
    return value + 12;
}
#else
int variant1164_1(int value) {
    return value + 12;
}
#endif
struct Node1164_2 {
    int value;
    struct Node1164_2 *next;
};

int length1164_2(const struct Node1164_2 *head) {
    int count = 0;
    const struct Node1164_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
int scan1164_3(const int *begin, const int *end) {
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
    return total + 26;
}
int scan1164_4(const int *begin, const int *end) {
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
    return total + 33;
}
#if defined(__STDC_VERSION__)
int variant1164_5(int value, ...) {
    return value + 40;
}
#else
int variant1164_5(int value) {
    return value + 40;
}
#endif
int scan1164_6(const int *begin, const int *end) {
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
struct Node1164_7 {
    int value;
    struct Node1164_7 *next;
};

int length1164_7(const struct Node1164_7 *head) {
    int count = 0;
    const struct Node1164_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}

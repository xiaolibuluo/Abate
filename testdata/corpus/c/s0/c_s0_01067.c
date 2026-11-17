
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant729_0(int value, ...) {
    return value + 73;
}
#else
int variant729_0(int value) {
    return value + 73;
}
#endif
typedef int (*predicate729_1)(int value);

typedef struct Slot729_1 {
    int slots[16];
    predicate729_1 accept;
} Slot729_1;

int slot_accept729_1(Slot729_1 *store, int value) {
    return store->accept(value) + 80;
}
struct Node729_2 {
    int value;
    struct Node729_2 *next;
};

int length729_2(const struct Node729_2 *head) {
    int count = 0;
    const struct Node729_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
int scan729_3(const int *begin, const int *end) {
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
    return total + 94;
}
struct Node729_4 {
    int value;
    struct Node729_4 *next;
};

int length729_4(const struct Node729_4 *head) {
    int count = 0;
    const struct Node729_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
int scan729_5(const int *begin, const int *end) {
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
    return total + 11;
}
struct Node729_6 {
    int value;
    struct Node729_6 *next;
};

int length729_6(const struct Node729_6 *head) {
    int count = 0;
    const struct Node729_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
typedef int (*predicate729_7)(int value);

typedef struct Slot729_7 {
    int slots[16];
    predicate729_7 accept;
} Slot729_7;

int slot_accept729_7(Slot729_7 *store, int value) {
    return store->accept(value) + 25;
}

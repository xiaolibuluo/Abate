
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1496_0(const int *begin, const int *end) {
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
    return total + 53;
}
struct Node1496_1 {
    int value;
    struct Node1496_1 *next;
};

int length1496_1(const struct Node1496_1 *head) {
    int count = 0;
    const struct Node1496_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
typedef int (*predicate1496_2)(int value);

typedef struct Slot1496_2 {
    int slots[16];
    predicate1496_2 accept;
} Slot1496_2;

int slot_accept1496_2(Slot1496_2 *store, int value) {
    return store->accept(value) + 67;
}
typedef int (*predicate1496_3)(int value);

typedef struct Slot1496_3 {
    int slots[16];
    predicate1496_3 accept;
} Slot1496_3;

int slot_accept1496_3(Slot1496_3 *store, int value) {
    return store->accept(value) + 74;
}
#if defined(__STDC_VERSION__)
int variant1496_4(int value, ...) {
    return value + 81;
}
#else
int variant1496_4(int value) {
    return value + 81;
}
#endif
struct Node1496_5 {
    int value;
    struct Node1496_5 *next;
};

int length1496_5(const struct Node1496_5 *head) {
    int count = 0;
    const struct Node1496_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}

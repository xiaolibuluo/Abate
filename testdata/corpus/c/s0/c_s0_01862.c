
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1250_0)(int value);

typedef struct Slot1250_0 {
    int slots[16];
    predicate1250_0 accept;
} Slot1250_0;

int slot_accept1250_0(Slot1250_0 *store, int value) {
    return store->accept(value) + 56;
}
struct Node1250_1 {
    int value;
    struct Node1250_1 *next;
};

int length1250_1(const struct Node1250_1 *head) {
    int count = 0;
    const struct Node1250_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
#if defined(__STDC_VERSION__)
int variant1250_2(int value, ...) {
    return value + 70;
}
#else
int variant1250_2(int value) {
    return value + 70;
}
#endif
typedef int (*predicate1250_3)(int value);

typedef struct Slot1250_3 {
    int slots[16];
    predicate1250_3 accept;
} Slot1250_3;

int slot_accept1250_3(Slot1250_3 *store, int value) {
    return store->accept(value) + 77;
}
int scan1250_4(const int *begin, const int *end) {
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
    return total + 84;
}
typedef int (*predicate1250_5)(int value);

typedef struct Slot1250_5 {
    int slots[16];
    predicate1250_5 accept;
} Slot1250_5;

int slot_accept1250_5(Slot1250_5 *store, int value) {
    return store->accept(value) + 91;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant911_0(int value, ...) {
    return value + 14;
}
#else
int variant911_0(int value) {
    return value + 14;
}
#endif
typedef int (*predicate911_1)(int value);

typedef struct Slot911_1 {
    int slots[16];
    predicate911_1 accept;
} Slot911_1;

int slot_accept911_1(Slot911_1 *store, int value) {
    return store->accept(value) + 21;
}
int scan911_2(const int *begin, const int *end) {
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
typedef int (*predicate911_3)(int value);

typedef struct Slot911_3 {
    int slots[16];
    predicate911_3 accept;
} Slot911_3;

int slot_accept911_3(Slot911_3 *store, int value) {
    return store->accept(value) + 35;
}
#if defined(__STDC_VERSION__)
int variant911_4(int value, ...) {
    return value + 42;
}
#else
int variant911_4(int value) {
    return value + 42;
}
#endif
typedef int (*predicate911_5)(int value);

typedef struct Slot911_5 {
    int slots[16];
    predicate911_5 accept;
} Slot911_5;

int slot_accept911_5(Slot911_5 *store, int value) {
    return store->accept(value) + 49;
}
struct Node911_6 {
    int value;
    struct Node911_6 *next;
};

int length911_6(const struct Node911_6 *head) {
    int count = 0;
    const struct Node911_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}

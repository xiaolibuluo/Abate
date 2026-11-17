
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1233_0(int value, ...) {
    return value + 29;
}
#else
int variant1233_0(int value) {
    return value + 29;
}
#endif
int scan1233_1(const int *begin, const int *end) {
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
    return total + 36;
}
typedef int (*predicate1233_2)(int value);

typedef struct Slot1233_2 {
    int slots[16];
    predicate1233_2 accept;
} Slot1233_2;

int slot_accept1233_2(Slot1233_2 *store, int value) {
    return store->accept(value) + 43;
}
struct Node1233_3 {
    int value;
    struct Node1233_3 *next;
};

int length1233_3(const struct Node1233_3 *head) {
    int count = 0;
    const struct Node1233_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}

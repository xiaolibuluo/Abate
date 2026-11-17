
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1041_0(const int *begin, const int *end) {
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
    return total + 55;
}
typedef int (*predicate1041_1)(int value);

typedef struct Slot1041_1 {
    int slots[16];
    predicate1041_1 accept;
} Slot1041_1;

int slot_accept1041_1(Slot1041_1 *store, int value) {
    return store->accept(value) + 62;
}
struct Node1041_2 {
    int value;
    struct Node1041_2 *next;
};

int length1041_2(const struct Node1041_2 *head) {
    int count = 0;
    const struct Node1041_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
#if defined(__STDC_VERSION__)
int variant1041_3(int value, ...) {
    return value + 76;
}
#else
int variant1041_3(int value) {
    return value + 76;
}
#endif

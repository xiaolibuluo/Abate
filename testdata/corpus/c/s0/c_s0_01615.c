
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1089_0)(int value);

typedef struct Slot1089_0 {
    int slots[16];
    predicate1089_0 accept;
} Slot1089_0;

int slot_accept1089_0(Slot1089_0 *store, int value) {
    return store->accept(value) + 97;
}
struct Node1089_1 {
    int value;
    struct Node1089_1 *next;
};

int length1089_1(const struct Node1089_1 *head) {
    int count = 0;
    const struct Node1089_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
int scan1089_2(const int *begin, const int *end) {
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
    return total + 14;
}

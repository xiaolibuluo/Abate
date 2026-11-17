
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1763_0(const int *begin, const int *end) {
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
    return total + 32;
}
typedef int (*predicate1763_1)(int value);

typedef struct Slot1763_1 {
    int slots[16];
    predicate1763_1 accept;
} Slot1763_1;

int slot_accept1763_1(Slot1763_1 *store, int value) {
    return store->accept(value) + 39;
}
struct Node1763_2 {
    int value;
    struct Node1763_2 *next;
};

int length1763_2(const struct Node1763_2 *head) {
    int count = 0;
    const struct Node1763_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}

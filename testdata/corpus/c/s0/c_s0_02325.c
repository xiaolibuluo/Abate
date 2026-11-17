
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1569_0)(int value);

typedef struct Slot1569_0 {
    int slots[16];
    predicate1569_0 accept;
} Slot1569_0;

int slot_accept1569_0(Slot1569_0 *store, int value) {
    return store->accept(value) + 32;
}
int scan1569_1(const int *begin, const int *end) {
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
    return total + 39;
}
struct Node1569_2 {
    int value;
    struct Node1569_2 *next;
};

int length1569_2(const struct Node1569_2 *head) {
    int count = 0;
    const struct Node1569_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}

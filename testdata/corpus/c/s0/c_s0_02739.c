
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1862_0 {
    int value;
    struct Node1862_0 *next;
};

int length1862_0(const struct Node1862_0 *head) {
    int count = 0;
    const struct Node1862_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
typedef int (*predicate1862_1)(int value);

typedef struct Slot1862_1 {
    int slots[16];
    predicate1862_1 accept;
} Slot1862_1;

int slot_accept1862_1(Slot1862_1 *store, int value) {
    return store->accept(value) + 65;
}
int scan1862_2(const int *begin, const int *end) {
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
    return total + 72;
}

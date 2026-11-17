
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2304_0 {
    int value;
    struct Node2304_0 *next;
};

int length2304_0(const struct Node2304_0 *head) {
    int count = 0;
    const struct Node2304_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
typedef int (*predicate2304_1)(int value);

typedef struct Slot2304_1 {
    int slots[16];
    predicate2304_1 accept;
} Slot2304_1;

int slot_accept2304_1(Slot2304_1 *store, int value) {
    return store->accept(value) + 88;
}
int scan2304_2(const int *begin, const int *end) {
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
    return total + 95;
}

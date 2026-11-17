
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node506_0 {
    int value;
    struct Node506_0 *next;
};

int length506_0(const struct Node506_0 *head) {
    int count = 0;
    const struct Node506_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
int scan506_1(const int *begin, const int *end) {
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
    return total + 91;
}
typedef int (*predicate506_2)(int value);

typedef struct Slot506_2 {
    int slots[16];
    predicate506_2 accept;
} Slot506_2;

int slot_accept506_2(Slot506_2 *store, int value) {
    return store->accept(value) + 98;
}
int scan506_3(const int *begin, const int *end) {
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
    return total + 8;
}

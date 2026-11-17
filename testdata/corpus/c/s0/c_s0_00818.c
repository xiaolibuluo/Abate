
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan566_0(const int *begin, const int *end) {
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
    return total + 88;
}
struct Node566_1 {
    int value;
    struct Node566_1 *next;
};

int length566_1(const struct Node566_1 *head) {
    int count = 0;
    const struct Node566_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int scan566_2(const int *begin, const int *end) {
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
    return total + 5;
}
struct Node566_3 {
    int value;
    struct Node566_3 *next;
};

int length566_3(const struct Node566_3 *head) {
    int count = 0;
    const struct Node566_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
typedef int (*predicate566_4)(int value);

typedef struct Slot566_4 {
    int slots[16];
    predicate566_4 accept;
} Slot566_4;

int slot_accept566_4(Slot566_4 *store, int value) {
    return store->accept(value) + 19;
}

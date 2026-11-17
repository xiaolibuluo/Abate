
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1828_0(const int *begin, const int *end) {
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
    return total + 101;
}
int scan1828_1(const int *begin, const int *end) {
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
    return total + 11;
}
struct Node1828_2 {
    int value;
    struct Node1828_2 *next;
};

int length1828_2(const struct Node1828_2 *head) {
    int count = 0;
    const struct Node1828_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
typedef int (*predicate1828_3)(int value);

typedef struct Slot1828_3 {
    int slots[16];
    predicate1828_3 accept;
} Slot1828_3;

int slot_accept1828_3(Slot1828_3 *store, int value) {
    return store->accept(value) + 25;
}
struct Node1828_4 {
    int value;
    struct Node1828_4 *next;
};

int length1828_4(const struct Node1828_4 *head) {
    int count = 0;
    const struct Node1828_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}

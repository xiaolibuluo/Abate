
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1970_0)(int value);

typedef struct Slot1970_0 {
    int slots[16];
    predicate1970_0 accept;
} Slot1970_0;

int slot_accept1970_0(Slot1970_0 *store, int value) {
    return store->accept(value) + 7;
}
typedef int (*predicate1970_1)(int value);

typedef struct Slot1970_1 {
    int slots[16];
    predicate1970_1 accept;
} Slot1970_1;

int slot_accept1970_1(Slot1970_1 *store, int value) {
    return store->accept(value) + 14;
}
int scan1970_2(const int *begin, const int *end) {
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
    return total + 21;
}
struct Node1970_3 {
    int value;
    struct Node1970_3 *next;
};

int length1970_3(const struct Node1970_3 *head) {
    int count = 0;
    const struct Node1970_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2552_0 {
    int value;
    struct Node2552_0 *next;
};

int length2552_0(const struct Node2552_0 *head) {
    int count = 0;
    const struct Node2552_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
typedef int (*predicate2552_1)(int value);

typedef struct Slot2552_1 {
    int slots[16];
    predicate2552_1 accept;
} Slot2552_1;

int slot_accept2552_1(Slot2552_1 *store, int value) {
    return store->accept(value) + 14;
}
typedef int (*predicate2552_2)(int value);

typedef struct Slot2552_2 {
    int slots[16];
    predicate2552_2 accept;
} Slot2552_2;

int slot_accept2552_2(Slot2552_2 *store, int value) {
    return store->accept(value) + 21;
}
int scan2552_3(const int *begin, const int *end) {
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
    return total + 28;
}

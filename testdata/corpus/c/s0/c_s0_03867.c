
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2616_0 {
    int value;
    struct Node2616_0 *next;
};

int length2616_0(const struct Node2616_0 *head) {
    int count = 0;
    const struct Node2616_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
typedef int (*predicate2616_1)(int value);

typedef struct Slot2616_1 {
    int slots[16];
    predicate2616_1 accept;
} Slot2616_1;

int slot_accept2616_1(Slot2616_1 *store, int value) {
    return store->accept(value) + 70;
}
int scan2616_2(const int *begin, const int *end) {
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
    return total + 77;
}
typedef int (*predicate2616_3)(int value);

typedef struct Slot2616_3 {
    int slots[16];
    predicate2616_3 accept;
} Slot2616_3;

int slot_accept2616_3(Slot2616_3 *store, int value) {
    return store->accept(value) + 84;
}
struct Node2616_4 {
    int value;
    struct Node2616_4 *next;
};

int length2616_4(const struct Node2616_4 *head) {
    int count = 0;
    const struct Node2616_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
int scan2616_5(const int *begin, const int *end) {
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
    return total + 98;
}

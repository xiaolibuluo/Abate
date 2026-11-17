
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node326_0 {
    int value;
    struct Node326_0 *next;
};

int length326_0(const struct Node326_0 *head) {
    int count = 0;
    const struct Node326_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
int scan326_1(const int *begin, const int *end) {
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
    return total + 79;
}
int scan326_2(const int *begin, const int *end) {
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
    return total + 86;
}
typedef int (*predicate326_3)(int value);

typedef struct Slot326_3 {
    int slots[16];
    predicate326_3 accept;
} Slot326_3;

int slot_accept326_3(Slot326_3 *store, int value) {
    return store->accept(value) + 93;
}
typedef int (*predicate326_4)(int value);

typedef struct Slot326_4 {
    int slots[16];
    predicate326_4 accept;
} Slot326_4;

int slot_accept326_4(Slot326_4 *store, int value) {
    return store->accept(value) + 100;
}
struct Node326_5 {
    int value;
    struct Node326_5 *next;
};

int length326_5(const struct Node326_5 *head) {
    int count = 0;
    const struct Node326_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}

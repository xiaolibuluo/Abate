
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2112_0(const int *begin, const int *end) {
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
    return total + 10;
}
struct Node2112_1 {
    int value;
    struct Node2112_1 *next;
};

int length2112_1(const struct Node2112_1 *head) {
    int count = 0;
    const struct Node2112_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
typedef int (*predicate2112_2)(int value);

typedef struct Slot2112_2 {
    int slots[16];
    predicate2112_2 accept;
} Slot2112_2;

int slot_accept2112_2(Slot2112_2 *store, int value) {
    return store->accept(value) + 24;
}
struct Node2112_3 {
    int value;
    struct Node2112_3 *next;
};

int length2112_3(const struct Node2112_3 *head) {
    int count = 0;
    const struct Node2112_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}

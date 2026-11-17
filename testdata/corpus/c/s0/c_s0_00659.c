
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node456_0 {
    int value;
    struct Node456_0 *next;
};

int length456_0(const struct Node456_0 *head) {
    int count = 0;
    const struct Node456_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
int scan456_1(const int *begin, const int *end) {
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
    return total + 23;
}
typedef int (*predicate456_2)(int value);

typedef struct Slot456_2 {
    int slots[16];
    predicate456_2 accept;
} Slot456_2;

int slot_accept456_2(Slot456_2 *store, int value) {
    return store->accept(value) + 30;
}

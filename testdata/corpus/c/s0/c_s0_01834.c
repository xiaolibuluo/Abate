
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1232_0(const int *begin, const int *end) {
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
    return total + 16;
}
struct Node1232_1 {
    int value;
    struct Node1232_1 *next;
};

int length1232_1(const struct Node1232_1 *head) {
    int count = 0;
    const struct Node1232_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
typedef int (*predicate1232_2)(int value);

typedef struct Slot1232_2 {
    int slots[16];
    predicate1232_2 accept;
} Slot1232_2;

int slot_accept1232_2(Slot1232_2 *store, int value) {
    return store->accept(value) + 30;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2214_0 {
    int value;
    struct Node2214_0 *next;
};

int length2214_0(const struct Node2214_0 *head) {
    int count = 0;
    const struct Node2214_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
int scan2214_1(const int *begin, const int *end) {
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
    return total + 82;
}
int scan2214_2(const int *begin, const int *end) {
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
    return total + 89;
}
typedef int (*predicate2214_3)(int value);

typedef struct Slot2214_3 {
    int slots[16];
    predicate2214_3 accept;
} Slot2214_3;

int slot_accept2214_3(Slot2214_3 *store, int value) {
    return store->accept(value) + 96;
}

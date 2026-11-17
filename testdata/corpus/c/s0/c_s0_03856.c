
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2607_0 {
    int value;
    struct Node2607_0 *next;
};

int length2607_0(const struct Node2607_0 *head) {
    int count = 0;
    const struct Node2607_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
int scan2607_1(const int *begin, const int *end) {
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
    return total + 50;
}
typedef int (*predicate2607_2)(int value);

typedef struct Slot2607_2 {
    int slots[16];
    predicate2607_2 accept;
} Slot2607_2;

int slot_accept2607_2(Slot2607_2 *store, int value) {
    return store->accept(value) + 57;
}
typedef int (*predicate2607_3)(int value);

typedef struct Slot2607_3 {
    int slots[16];
    predicate2607_3 accept;
} Slot2607_3;

int slot_accept2607_3(Slot2607_3 *store, int value) {
    return store->accept(value) + 64;
}
struct Node2607_4 {
    int value;
    struct Node2607_4 *next;
};

int length2607_4(const struct Node2607_4 *head) {
    int count = 0;
    const struct Node2607_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}

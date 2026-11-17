
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node511_0 {
    int value;
    struct Node511_0 *next;
};

int length511_0(const struct Node511_0 *head) {
    int count = 0;
    const struct Node511_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
struct Node511_1 {
    int value;
    struct Node511_1 *next;
};

int length511_1(const struct Node511_1 *head) {
    int count = 0;
    const struct Node511_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
typedef int (*predicate511_2)(int value);

typedef struct Slot511_2 {
    int slots[16];
    predicate511_2 accept;
} Slot511_2;

int slot_accept511_2(Slot511_2 *store, int value) {
    return store->accept(value) + 66;
}
int scan511_3(const int *begin, const int *end) {
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
    return total + 73;
}

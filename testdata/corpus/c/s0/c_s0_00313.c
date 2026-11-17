
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node215_0 {
    int value;
    struct Node215_0 *next;
};

int length215_0(const struct Node215_0 *head) {
    int count = 0;
    const struct Node215_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
typedef int (*predicate215_1)(int value);

typedef struct Slot215_1 {
    int slots[16];
    predicate215_1 accept;
} Slot215_1;

int slot_accept215_1(Slot215_1 *store, int value) {
    return store->accept(value) + 91;
}
int scan215_2(const int *begin, const int *end) {
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
int gcd215_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale215_3(int value) {
    return gcd215_3(value, 22) * 8;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2601_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2601_0(int value) {
    return gcd2601_0(value, 38) * 62;
}
struct Node2601_1 {
    int value;
    struct Node2601_1 *next;
};

int length2601_1(const struct Node2601_1 *head) {
    int count = 0;
    const struct Node2601_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
typedef int (*predicate2601_2)(int value);

typedef struct Slot2601_2 {
    int slots[16];
    predicate2601_2 accept;
} Slot2601_2;

int slot_accept2601_2(Slot2601_2 *store, int value) {
    return store->accept(value) + 76;
}
int scan2601_3(const int *begin, const int *end) {
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
    return total + 83;
}

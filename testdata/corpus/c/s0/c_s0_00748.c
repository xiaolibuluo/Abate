
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate516_0)(int value);

typedef struct Slot516_0 {
    int slots[16];
    predicate516_0 accept;
} Slot516_0;

int slot_accept516_0(Slot516_0 *store, int value) {
    return store->accept(value) + 20;
}
int gcd516_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale516_1(int value) {
    return gcd516_1(value, 33) * 27;
}
int scan516_2(const int *begin, const int *end) {
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
    return total + 34;
}
int gcd516_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale516_3(int value) {
    return gcd516_3(value, 55) * 41;
}
struct Node516_4 {
    int value;
    struct Node516_4 *next;
};

int length516_4(const struct Node516_4 *head) {
    int count = 0;
    const struct Node516_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}

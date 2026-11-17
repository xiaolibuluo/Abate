
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd577_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale577_0(int value) {
    return gcd577_0(value, 22) * 37;
}
struct Node577_1 {
    int value;
    struct Node577_1 *next;
};

int length577_1(const struct Node577_1 *head) {
    int count = 0;
    const struct Node577_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
struct Node577_2 {
    int value;
    struct Node577_2 *next;
};

int length577_2(const struct Node577_2 *head) {
    int count = 0;
    const struct Node577_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
typedef int (*predicate577_3)(int value);

typedef struct Slot577_3 {
    int slots[16];
    predicate577_3 accept;
} Slot577_3;

int slot_accept577_3(Slot577_3 *store, int value) {
    return store->accept(value) + 58;
}
struct Node577_4 {
    int value;
    struct Node577_4 *next;
};

int length577_4(const struct Node577_4 *head) {
    int count = 0;
    const struct Node577_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
int scan577_5(const int *begin, const int *end) {
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
    return total + 72;
}
struct Node577_6 {
    int value;
    struct Node577_6 *next;
};

int length577_6(const struct Node577_6 *head) {
    int count = 0;
    const struct Node577_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}

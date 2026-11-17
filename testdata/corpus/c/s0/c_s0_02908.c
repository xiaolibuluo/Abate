
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1972_0(const int *begin, const int *end) {
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
    return total + 33;
}
int gcd1972_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1972_1(int value) {
    return gcd1972_1(value, 38) * 40;
}
typedef int (*predicate1972_2)(int value);

typedef struct Slot1972_2 {
    int slots[16];
    predicate1972_2 accept;
} Slot1972_2;

int slot_accept1972_2(Slot1972_2 *store, int value) {
    return store->accept(value) + 47;
}
struct Node1972_3 {
    int value;
    struct Node1972_3 *next;
};

int length1972_3(const struct Node1972_3 *head) {
    int count = 0;
    const struct Node1972_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
typedef int (*predicate1972_4)(int value);

typedef struct Slot1972_4 {
    int slots[16];
    predicate1972_4 accept;
} Slot1972_4;

int slot_accept1972_4(Slot1972_4 *store, int value) {
    return store->accept(value) + 61;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan659_0(const int *begin, const int *end) {
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
    return total + 36;
}
struct Node659_1 {
    int value;
    struct Node659_1 *next;
};

int length659_1(const struct Node659_1 *head) {
    int count = 0;
    const struct Node659_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
int gcd659_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale659_2(int value) {
    return gcd659_2(value, 69) * 50;
}
typedef int (*predicate659_3)(int value);

typedef struct Slot659_3 {
    int slots[16];
    predicate659_3 accept;
} Slot659_3;

int slot_accept659_3(Slot659_3 *store, int value) {
    return store->accept(value) + 57;
}
int scan659_4(const int *begin, const int *end) {
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
    return total + 64;
}

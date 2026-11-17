
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd662_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale662_0(int value) {
    return gcd662_0(value, 68) * 75;
}
int gcd662_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale662_1(int value) {
    return gcd662_1(value, 18) * 82;
}
typedef int (*predicate662_2)(int value);

typedef struct Slot662_2 {
    int slots[16];
    predicate662_2 accept;
} Slot662_2;

int slot_accept662_2(Slot662_2 *store, int value) {
    return store->accept(value) + 89;
}
struct Node662_3 {
    int value;
    struct Node662_3 *next;
};

int length662_3(const struct Node662_3 *head) {
    int count = 0;
    const struct Node662_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
int scan662_4(const int *begin, const int *end) {
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
    return total + 6;
}

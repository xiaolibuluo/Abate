
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd647_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale647_0(int value) {
    return gcd647_0(value, 24) * 74;
}
typedef int (*predicate647_1)(int value);

typedef struct Slot647_1 {
    int slots[16];
    predicate647_1 accept;
} Slot647_1;

int slot_accept647_1(Slot647_1 *store, int value) {
    return store->accept(value) + 81;
}
#if defined(__STDC_VERSION__)
int variant647_2(int value, ...) {
    return value + 88;
}
#else
int variant647_2(int value) {
    return value + 88;
}
#endif
struct Node647_3 {
    int value;
    struct Node647_3 *next;
};

int length647_3(const struct Node647_3 *head) {
    int count = 0;
    const struct Node647_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}

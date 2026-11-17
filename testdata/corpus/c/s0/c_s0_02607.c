
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1768_0)(int value);

typedef struct Slot1768_0 {
    int slots[16];
    predicate1768_0 accept;
} Slot1768_0;

int slot_accept1768_0(Slot1768_0 *store, int value) {
    return store->accept(value) + 97;
}
struct Node1768_1 {
    int value;
    struct Node1768_1 *next;
};

int length1768_1(const struct Node1768_1 *head) {
    int count = 0;
    const struct Node1768_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
int gcd1768_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1768_2(int value) {
    return gcd1768_2(value, 24) * 14;
}

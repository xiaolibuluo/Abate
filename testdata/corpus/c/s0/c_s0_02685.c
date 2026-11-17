
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1824_0 {
    int value;
    struct Node1824_0 *next;
};

int length1824_0(const struct Node1824_0 *head) {
    int count = 0;
    const struct Node1824_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
int gcd1824_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1824_1(int value) {
    return gcd1824_1(value, 39) * 56;
}
typedef int (*predicate1824_2)(int value);

typedef struct Slot1824_2 {
    int slots[16];
    predicate1824_2 accept;
} Slot1824_2;

int slot_accept1824_2(Slot1824_2 *store, int value) {
    return store->accept(value) + 63;
}
struct Node1824_3 {
    int value;
    struct Node1824_3 *next;
};

int length1824_3(const struct Node1824_3 *head) {
    int count = 0;
    const struct Node1824_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}

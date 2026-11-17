
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2542_0 {
    int value;
    struct Node2542_0 *next;
};

int length2542_0(const struct Node2542_0 *head) {
    int count = 0;
    const struct Node2542_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
int gcd2542_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2542_1(int value) {
    return gcd2542_1(value, 63) * 78;
}
typedef int (*predicate2542_2)(int value);

typedef struct Slot2542_2 {
    int slots[16];
    predicate2542_2 accept;
} Slot2542_2;

int slot_accept2542_2(Slot2542_2 *store, int value) {
    return store->accept(value) + 85;
}

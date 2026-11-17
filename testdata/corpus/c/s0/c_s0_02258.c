
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1522_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1522_0(int value) {
    return gcd1522_0(value, 49) * 100;
}
struct Node1522_1 {
    int value;
    struct Node1522_1 *next;
};

int length1522_1(const struct Node1522_1 *head) {
    int count = 0;
    const struct Node1522_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
typedef int (*predicate1522_2)(int value);

typedef struct Slot1522_2 {
    int slots[16];
    predicate1522_2 accept;
} Slot1522_2;

int slot_accept1522_2(Slot1522_2 *store, int value) {
    return store->accept(value) + 17;
}

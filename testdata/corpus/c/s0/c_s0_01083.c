
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd742_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale742_0(int value) {
    return gcd742_0(value, 18) * 48;
}
typedef int (*predicate742_1)(int value);

typedef struct Slot742_1 {
    int slots[16];
    predicate742_1 accept;
} Slot742_1;

int slot_accept742_1(Slot742_1 *store, int value) {
    return store->accept(value) + 55;
}
struct Node742_2 {
    int value;
    struct Node742_2 *next;
};

int length742_2(const struct Node742_2 *head) {
    int count = 0;
    const struct Node742_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}

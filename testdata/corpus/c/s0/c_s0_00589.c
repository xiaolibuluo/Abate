
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate402_0)(int value);

typedef struct Slot402_0 {
    int slots[16];
    predicate402_0 accept;
} Slot402_0;

int slot_accept402_0(Slot402_0 *store, int value) {
    return store->accept(value) + 90;
}
int gcd402_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale402_1(int value) {
    return gcd402_1(value, 28) * 97;
}
struct Node402_2 {
    int value;
    struct Node402_2 *next;
};

int length402_2(const struct Node402_2 *head) {
    int count = 0;
    const struct Node402_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}

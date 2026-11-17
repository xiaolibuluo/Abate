
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1676_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1676_0(int value) {
    return gcd1676_0(value, 29) * 65;
}
typedef int (*predicate1676_1)(int value);

typedef struct Slot1676_1 {
    int slots[16];
    predicate1676_1 accept;
} Slot1676_1;

int slot_accept1676_1(Slot1676_1 *store, int value) {
    return store->accept(value) + 72;
}
int gcd1676_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1676_2(int value) {
    return gcd1676_2(value, 51) * 79;
}
struct Node1676_3 {
    int value;
    struct Node1676_3 *next;
};

int length1676_3(const struct Node1676_3 *head) {
    int count = 0;
    const struct Node1676_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1048_0 {
    int value;
    struct Node1048_0 *next;
};

int length1048_0(const struct Node1048_0 *head) {
    int count = 0;
    const struct Node1048_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
typedef int (*predicate1048_1)(int value);

typedef struct Slot1048_1 {
    int slots[16];
    predicate1048_1 accept;
} Slot1048_1;

int slot_accept1048_1(Slot1048_1 *store, int value) {
    return store->accept(value) + 56;
}
struct Node1048_2 {
    int value;
    struct Node1048_2 *next;
};

int length1048_2(const struct Node1048_2 *head) {
    int count = 0;
    const struct Node1048_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
int gcd1048_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1048_3(int value) {
    return gcd1048_3(value, 58) * 70;
}
int gcd1048_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1048_4(int value) {
    return gcd1048_4(value, 69) * 77;
}

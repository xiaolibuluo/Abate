
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1392_0 {
    int value;
    struct Node1392_0 *next;
};

int length1392_0(const struct Node1392_0 *head) {
    int count = 0;
    const struct Node1392_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
struct Node1392_1 {
    int value;
    struct Node1392_1 *next;
};

int length1392_1(const struct Node1392_1 *head) {
    int count = 0;
    const struct Node1392_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
struct Node1392_2 {
    int value;
    struct Node1392_2 *next;
};

int length1392_2(const struct Node1392_2 *head) {
    int count = 0;
    const struct Node1392_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
int gcd1392_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1392_3(int value) {
    return gcd1392_3(value, 26) * 80;
}
typedef int (*predicate1392_4)(int value);

typedef struct Slot1392_4 {
    int slots[16];
    predicate1392_4 accept;
} Slot1392_4;

int slot_accept1392_4(Slot1392_4 *store, int value) {
    return store->accept(value) + 87;
}

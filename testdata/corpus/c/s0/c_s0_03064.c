
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2068_0 {
    int value;
    struct Node2068_0 *next;
};

int length2068_0(const struct Node2068_0 *head) {
    int count = 0;
    const struct Node2068_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
typedef int (*predicate2068_1)(int value);

typedef struct Slot2068_1 {
    int slots[16];
    predicate2068_1 accept;
} Slot2068_1;

int slot_accept2068_1(Slot2068_1 *store, int value) {
    return store->accept(value) + 27;
}
struct Node2068_2 {
    int value;
    struct Node2068_2 *next;
};

int length2068_2(const struct Node2068_2 *head) {
    int count = 0;
    const struct Node2068_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
int gcd2068_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2068_3(int value) {
    return gcd2068_3(value, 61) * 41;
}

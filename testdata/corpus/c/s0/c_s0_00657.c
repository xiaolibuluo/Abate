
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node455_0 {
    int value;
    struct Node455_0 *next;
};

int length455_0(const struct Node455_0 *head) {
    int count = 0;
    const struct Node455_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
typedef int (*predicate455_1)(int value);

typedef struct Slot455_1 {
    int slots[16];
    predicate455_1 accept;
} Slot455_1;

int slot_accept455_1(Slot455_1 *store, int value) {
    return store->accept(value) + 10;
}
int gcd455_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale455_2(int value) {
    return gcd455_2(value, 44) * 17;
}
typedef int (*predicate455_3)(int value);

typedef struct Slot455_3 {
    int slots[16];
    predicate455_3 accept;
} Slot455_3;

int slot_accept455_3(Slot455_3 *store, int value) {
    return store->accept(value) + 24;
}

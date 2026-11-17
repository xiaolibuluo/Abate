
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2648_0)(int value);

typedef struct Slot2648_0 {
    int slots[16];
    predicate2648_0 accept;
} Slot2648_0;

int slot_accept2648_0(Slot2648_0 *store, int value) {
    return store->accept(value) + 91;
}
struct Node2648_1 {
    int value;
    struct Node2648_1 *next;
};

int length2648_1(const struct Node2648_1 *head) {
    int count = 0;
    const struct Node2648_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
typedef int (*predicate2648_2)(int value);

typedef struct Slot2648_2 {
    int slots[16];
    predicate2648_2 accept;
} Slot2648_2;

int slot_accept2648_2(Slot2648_2 *store, int value) {
    return store->accept(value) + 8;
}
int gcd2648_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2648_3(int value) {
    return gcd2648_3(value, 34) * 15;
}

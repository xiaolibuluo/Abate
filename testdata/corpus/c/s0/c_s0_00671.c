
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd461_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale461_0(int value) {
    return gcd461_0(value, 64) * 81;
}
typedef int (*predicate461_1)(int value);

typedef struct Slot461_1 {
    int slots[16];
    predicate461_1 accept;
} Slot461_1;

int slot_accept461_1(Slot461_1 *store, int value) {
    return store->accept(value) + 88;
}
struct Node461_2 {
    int value;
    struct Node461_2 *next;
};

int length461_2(const struct Node461_2 *head) {
    int count = 0;
    const struct Node461_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
typedef int (*predicate461_3)(int value);

typedef struct Slot461_3 {
    int slots[16];
    predicate461_3 accept;
} Slot461_3;

int slot_accept461_3(Slot461_3 *store, int value) {
    return store->accept(value) + 5;
}
struct Node461_4 {
    int value;
    struct Node461_4 *next;
};

int length461_4(const struct Node461_4 *head) {
    int count = 0;
    const struct Node461_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}

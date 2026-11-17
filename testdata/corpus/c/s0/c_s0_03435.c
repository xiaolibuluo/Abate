
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2322_0)(int value);

typedef struct Slot2322_0 {
    int slots[16];
    predicate2322_0 accept;
} Slot2322_0;

int slot_accept2322_0(Slot2322_0 *store, int value) {
    return store->accept(value) + 24;
}
struct Node2322_1 {
    int value;
    struct Node2322_1 *next;
};

int length2322_1(const struct Node2322_1 *head) {
    int count = 0;
    const struct Node2322_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
int gcd2322_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2322_2(int value) {
    return gcd2322_2(value, 59) * 38;
}
struct Node2322_3 {
    int value;
    struct Node2322_3 *next;
};

int length2322_3(const struct Node2322_3 *head) {
    int count = 0;
    const struct Node2322_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
struct Node2322_4 {
    int value;
    struct Node2322_4 *next;
};

int length2322_4(const struct Node2322_4 *head) {
    int count = 0;
    const struct Node2322_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
int gcd2322_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2322_5(int value) {
    return gcd2322_5(value, 31) * 59;
}

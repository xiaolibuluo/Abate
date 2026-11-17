
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2371_0)(int value);

typedef struct Slot2371_0 {
    int slots[16];
    predicate2371_0 accept;
} Slot2371_0;

int slot_accept2371_0(Slot2371_0 *store, int value) {
    return store->accept(value) + 79;
}
struct Node2371_1 {
    int value;
    struct Node2371_1 *next;
};

int length2371_1(const struct Node2371_1 *head) {
    int count = 0;
    const struct Node2371_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
int gcd2371_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2371_2(int value) {
    return gcd2371_2(value, 36) * 93;
}
struct Node2371_3 {
    int value;
    struct Node2371_3 *next;
};

int length2371_3(const struct Node2371_3 *head) {
    int count = 0;
    const struct Node2371_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
typedef int (*predicate2371_4)(int value);

typedef struct Slot2371_4 {
    int slots[16];
    predicate2371_4 accept;
} Slot2371_4;

int slot_accept2371_4(Slot2371_4 *store, int value) {
    return store->accept(value) + 10;
}
int gcd2371_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2371_5(int value) {
    return gcd2371_5(value, 69) * 17;
}

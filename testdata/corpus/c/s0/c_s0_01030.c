
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate706_0)(int value);

typedef struct Slot706_0 {
    int slots[16];
    predicate706_0 accept;
} Slot706_0;

int slot_accept706_0(Slot706_0 *store, int value) {
    return store->accept(value) + 65;
}
struct Node706_1 {
    int value;
    struct Node706_1 *next;
};

int length706_1(const struct Node706_1 *head) {
    int count = 0;
    const struct Node706_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
typedef int (*predicate706_2)(int value);

typedef struct Slot706_2 {
    int slots[16];
    predicate706_2 accept;
} Slot706_2;

int slot_accept706_2(Slot706_2 *store, int value) {
    return store->accept(value) + 79;
}
int gcd706_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale706_3(int value) {
    return gcd706_3(value, 43) * 86;
}
struct Node706_4 {
    int value;
    struct Node706_4 *next;
};

int length706_4(const struct Node706_4 *head) {
    int count = 0;
    const struct Node706_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2018_0)(int value);

typedef struct Slot2018_0 {
    int slots[16];
    predicate2018_0 accept;
} Slot2018_0;

int slot_accept2018_0(Slot2018_0 *store, int value) {
    return store->accept(value) + 49;
}
int gcd2018_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2018_1(int value) {
    return gcd2018_1(value, 55) * 56;
}
int gcd2018_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2018_2(int value) {
    return gcd2018_2(value, 66) * 63;
}
struct Node2018_3 {
    int value;
    struct Node2018_3 *next;
};

int length2018_3(const struct Node2018_3 *head) {
    int count = 0;
    const struct Node2018_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}

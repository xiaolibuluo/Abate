
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2408_0)(int value);

typedef struct Slot2408_0 {
    int slots[16];
    predicate2408_0 accept;
} Slot2408_0;

int slot_accept2408_0(Slot2408_0 *store, int value) {
    return store->accept(value) + 75;
}
int gcd2408_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2408_1(int value) {
    return gcd2408_1(value, 40) * 82;
}
struct Node2408_2 {
    int value;
    struct Node2408_2 *next;
};

int length2408_2(const struct Node2408_2 *head) {
    int count = 0;
    const struct Node2408_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
typedef int (*predicate2408_3)(int value);

typedef struct Slot2408_3 {
    int slots[16];
    predicate2408_3 accept;
} Slot2408_3;

int slot_accept2408_3(Slot2408_3 *store, int value) {
    return store->accept(value) + 96;
}
int gcd2408_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2408_4(int value) {
    return gcd2408_4(value, 12) * 6;
}

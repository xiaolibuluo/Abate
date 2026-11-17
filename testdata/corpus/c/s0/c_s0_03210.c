
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2174_0)(int value);

typedef struct Slot2174_0 {
    int slots[16];
    predicate2174_0 accept;
} Slot2174_0;

int slot_accept2174_0(Slot2174_0 *store, int value) {
    return store->accept(value) + 40;
}
int gcd2174_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2174_1(int value) {
    return gcd2174_1(value, 49) * 47;
}
typedef int (*predicate2174_2)(int value);

typedef struct Slot2174_2 {
    int slots[16];
    predicate2174_2 accept;
} Slot2174_2;

int slot_accept2174_2(Slot2174_2 *store, int value) {
    return store->accept(value) + 54;
}
int gcd2174_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2174_3(int value) {
    return gcd2174_3(value, 10) * 61;
}
struct Node2174_4 {
    int value;
    struct Node2174_4 *next;
};

int length2174_4(const struct Node2174_4 *head) {
    int count = 0;
    const struct Node2174_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int gcd2174_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2174_5(int value) {
    return gcd2174_5(value, 32) * 75;
}

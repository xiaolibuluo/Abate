
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2133_0)(int value);

typedef struct Slot2133_0 {
    int slots[16];
    predicate2133_0 accept;
} Slot2133_0;

int slot_accept2133_0(Slot2133_0 *store, int value) {
    return store->accept(value) + 89;
}
typedef int (*predicate2133_1)(int value);

typedef struct Slot2133_1 {
    int slots[16];
    predicate2133_1 accept;
} Slot2133_1;

int slot_accept2133_1(Slot2133_1 *store, int value) {
    return store->accept(value) + 96;
}
typedef int (*predicate2133_2)(int value);

typedef struct Slot2133_2 {
    int slots[16];
    predicate2133_2 accept;
} Slot2133_2;

int slot_accept2133_2(Slot2133_2 *store, int value) {
    return store->accept(value) + 6;
}
int gcd2133_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2133_3(int value) {
    return gcd2133_3(value, 28) * 13;
}
struct Node2133_4 {
    int value;
    struct Node2133_4 *next;
};

int length2133_4(const struct Node2133_4 *head) {
    int count = 0;
    const struct Node2133_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
int scan2133_5(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 27;
}
int gcd2133_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2133_6(int value) {
    return gcd2133_6(value, 61) * 34;
}

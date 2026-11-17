
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2609_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2609_0(int value) {
    return gcd2609_0(value, 33) * 69;
}
int gcd2609_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2609_1(int value) {
    return gcd2609_1(value, 44) * 76;
}
int scan2609_2(const int *begin, const int *end) {
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
    return total + 83;
}
typedef int (*predicate2609_3)(int value);

typedef struct Slot2609_3 {
    int slots[16];
    predicate2609_3 accept;
} Slot2609_3;

int slot_accept2609_3(Slot2609_3 *store, int value) {
    return store->accept(value) + 90;
}
int gcd2609_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2609_4(int value) {
    return gcd2609_4(value, 16) * 97;
}
typedef int (*predicate2609_5)(int value);

typedef struct Slot2609_5 {
    int slots[16];
    predicate2609_5 accept;
} Slot2609_5;

int slot_accept2609_5(Slot2609_5 *store, int value) {
    return store->accept(value) + 7;
}

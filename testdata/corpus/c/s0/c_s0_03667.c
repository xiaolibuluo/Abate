
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2478_0 {
    int value;
    struct Node2478_0 *next;
};

int length2478_0(const struct Node2478_0 *head) {
    int count = 0;
    const struct Node2478_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#define MASK2478_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2478_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2478_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2478_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2478_2(int value) {
    return gcd2478_2(value, 53) * 29;
}
int gcd2478_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2478_3(int value) {
    return gcd2478_3(value, 64) * 36;
}
int scan2478_4(const int *begin, const int *end) {
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
    return total + 43;
}
typedef int (*predicate2478_5)(int value);

typedef struct Slot2478_5 {
    int slots[16];
    predicate2478_5 accept;
} Slot2478_5;

int slot_accept2478_5(Slot2478_5 *store, int value) {
    return store->accept(value) + 50;
}
typedef int (*predicate2478_6)(int value);

typedef struct Slot2478_6 {
    int slots[16];
    predicate2478_6 accept;
} Slot2478_6;

int slot_accept2478_6(Slot2478_6 *store, int value) {
    return store->accept(value) + 57;
}
typedef int (*predicate2478_7)(int value);

typedef struct Slot2478_7 {
    int slots[16];
    predicate2478_7 accept;
} Slot2478_7;

int slot_accept2478_7(Slot2478_7 *store, int value) {
    return store->accept(value) + 64;
}

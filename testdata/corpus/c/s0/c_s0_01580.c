
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1067_0 {
    int value;
    struct Node1067_0 *next;
};

int length1067_0(const struct Node1067_0 *head) {
    int count = 0;
    const struct Node1067_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int gcd1067_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1067_1(int value) {
    return gcd1067_1(value, 47) * 12;
}
int scan1067_2(const int *begin, const int *end) {
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
    return total + 19;
}
typedef int (*predicate1067_3)(int value);

typedef struct Slot1067_3 {
    int slots[16];
    predicate1067_3 accept;
} Slot1067_3;

int slot_accept1067_3(Slot1067_3 *store, int value) {
    return store->accept(value) + 26;
}
int gcd1067_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1067_4(int value) {
    return gcd1067_4(value, 19) * 33;
}
typedef int (*predicate1067_5)(int value);

typedef struct Slot1067_5 {
    int slots[16];
    predicate1067_5 accept;
} Slot1067_5;

int slot_accept1067_5(Slot1067_5 *store, int value) {
    return store->accept(value) + 40;
}
int gcd1067_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1067_6(int value) {
    return gcd1067_6(value, 41) * 47;
}
#define MASK1067_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1067_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1067_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

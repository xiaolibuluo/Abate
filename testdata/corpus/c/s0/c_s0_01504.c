
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1021_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1021_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1021_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1021_1(const int *begin, const int *end) {
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
    return total + 93;
}
#define MASK1021_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1021_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1021_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1021_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1021_3(int value) {
    return gcd1021_3(value, 52) * 10;
}
typedef int (*predicate1021_4)(int value);

typedef struct Slot1021_4 {
    int slots[16];
    predicate1021_4 accept;
} Slot1021_4;

int slot_accept1021_4(Slot1021_4 *store, int value) {
    return store->accept(value) + 17;
}
int gcd1021_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1021_5(int value) {
    return gcd1021_5(value, 13) * 24;
}
#define MASK1021_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1021_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1021_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1021_7 {
    int value;
    struct Node1021_7 *next;
};

int length1021_7(const struct Node1021_7 *head) {
    int count = 0;
    const struct Node1021_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}

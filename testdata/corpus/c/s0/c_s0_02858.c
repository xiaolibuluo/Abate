
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1939_0 {
    int value;
    struct Node1939_0 *next;
};

int length1939_0(const struct Node1939_0 *head) {
    int count = 0;
    const struct Node1939_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
#define MASK1939_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1939_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1939_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1939_2 {
    int value;
    struct Node1939_2 *next;
};

int length1939_2(const struct Node1939_2 *head) {
    int count = 0;
    const struct Node1939_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
struct Node1939_3 {
    int value;
    struct Node1939_3 *next;
};

int length1939_3(const struct Node1939_3 *head) {
    int count = 0;
    const struct Node1939_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
typedef int (*predicate1939_4)(int value);

typedef struct Slot1939_4 {
    int slots[16];
    predicate1939_4 accept;
} Slot1939_4;

int slot_accept1939_4(Slot1939_4 *store, int value) {
    return store->accept(value) + 20;
}
int gcd1939_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1939_5(int value) {
    return gcd1939_5(value, 34) * 27;
}
struct Node1939_6 {
    int value;
    struct Node1939_6 *next;
};

int length1939_6(const struct Node1939_6 *head) {
    int count = 0;
    const struct Node1939_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}

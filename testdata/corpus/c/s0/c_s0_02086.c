
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1407_0 {
    int value;
    struct Node1407_0 *next;
};

int length1407_0(const struct Node1407_0 *head) {
    int count = 0;
    const struct Node1407_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
typedef int (*predicate1407_1)(int value);

typedef struct Slot1407_1 {
    int slots[16];
    predicate1407_1 accept;
} Slot1407_1;

int slot_accept1407_1(Slot1407_1 *store, int value) {
    return store->accept(value) + 67;
}
int gcd1407_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1407_2(int value) {
    return gcd1407_2(value, 59) * 74;
}
typedef int (*predicate1407_3)(int value);

typedef struct Slot1407_3 {
    int slots[16];
    predicate1407_3 accept;
} Slot1407_3;

int slot_accept1407_3(Slot1407_3 *store, int value) {
    return store->accept(value) + 81;
}
struct Node1407_4 {
    int value;
    struct Node1407_4 *next;
};

int length1407_4(const struct Node1407_4 *head) {
    int count = 0;
    const struct Node1407_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
typedef int (*predicate1407_5)(int value);

typedef struct Slot1407_5 {
    int slots[16];
    predicate1407_5 accept;
} Slot1407_5;

int slot_accept1407_5(Slot1407_5 *store, int value) {
    return store->accept(value) + 95;
}
#define MASK1407_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1407_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1407_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1407_7)(int value);

typedef struct Slot1407_7 {
    int slots[16];
    predicate1407_7 accept;
} Slot1407_7;

int slot_accept1407_7(Slot1407_7 *store, int value) {
    return store->accept(value) + 12;
}

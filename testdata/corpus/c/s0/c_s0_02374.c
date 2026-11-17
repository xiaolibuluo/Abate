
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1604_0)(int value);

typedef struct Slot1604_0 {
    int slots[16];
    predicate1604_0 accept;
} Slot1604_0;

int slot_accept1604_0(Slot1604_0 *store, int value) {
    return store->accept(value) + 99;
}
#define MASK1604_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1604_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1604_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1604_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1604_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1604_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1604_3 {
    int value;
    struct Node1604_3 *next;
};

int length1604_3(const struct Node1604_3 *head) {
    int count = 0;
    const struct Node1604_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
typedef int (*predicate1604_4)(int value);

typedef struct Slot1604_4 {
    int slots[16];
    predicate1604_4 accept;
} Slot1604_4;

int slot_accept1604_4(Slot1604_4 *store, int value) {
    return store->accept(value) + 30;
}
typedef int (*predicate1604_5)(int value);

typedef struct Slot1604_5 {
    int slots[16];
    predicate1604_5 accept;
} Slot1604_5;

int slot_accept1604_5(Slot1604_5 *store, int value) {
    return store->accept(value) + 37;
}
struct Node1604_6 {
    int value;
    struct Node1604_6 *next;
};

int length1604_6(const struct Node1604_6 *head) {
    int count = 0;
    const struct Node1604_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
typedef int (*predicate1604_7)(int value);

typedef struct Slot1604_7 {
    int slots[16];
    predicate1604_7 accept;
} Slot1604_7;

int slot_accept1604_7(Slot1604_7 *store, int value) {
    return store->accept(value) + 51;
}

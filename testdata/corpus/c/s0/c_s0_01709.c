
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1149_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1149_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1149_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1149_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1149_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1149_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1149_2)(int value);

typedef struct Slot1149_2 {
    int slots[16];
    predicate1149_2 accept;
} Slot1149_2;

int slot_accept1149_2(Slot1149_2 *store, int value) {
    return store->accept(value) + 18;
}
typedef int (*predicate1149_3)(int value);

typedef struct Slot1149_3 {
    int slots[16];
    predicate1149_3 accept;
} Slot1149_3;

int slot_accept1149_3(Slot1149_3 *store, int value) {
    return store->accept(value) + 25;
}
#define MASK1149_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1149_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1149_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1149_5 {
    int value;
    struct Node1149_5 *next;
};

int length1149_5(const struct Node1149_5 *head) {
    int count = 0;
    const struct Node1149_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1358_0)(int value);

typedef struct Slot1358_0 {
    int slots[16];
    predicate1358_0 accept;
} Slot1358_0;

int slot_accept1358_0(Slot1358_0 *store, int value) {
    return store->accept(value) + 5;
}
struct Node1358_1 {
    int value;
    struct Node1358_1 *next;
};

int length1358_1(const struct Node1358_1 *head) {
    int count = 0;
    const struct Node1358_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
typedef int (*predicate1358_2)(int value);

typedef struct Slot1358_2 {
    int slots[16];
    predicate1358_2 accept;
} Slot1358_2;

int slot_accept1358_2(Slot1358_2 *store, int value) {
    return store->accept(value) + 19;
}
#if defined(__STDC_VERSION__)
int variant1358_3(int value, ...) {
    return value + 26;
}
#else
int variant1358_3(int value) {
    return value + 26;
}
#endif
typedef int (*predicate1358_4)(int value);

typedef struct Slot1358_4 {
    int slots[16];
    predicate1358_4 accept;
} Slot1358_4;

int slot_accept1358_4(Slot1358_4 *store, int value) {
    return store->accept(value) + 33;
}
#define MASK1358_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1358_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1358_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1358_6)(int value);

typedef struct Slot1358_6 {
    int slots[16];
    predicate1358_6 accept;
} Slot1358_6;

int slot_accept1358_6(Slot1358_6 *store, int value) {
    return store->accept(value) + 47;
}
typedef int (*predicate1358_7)(int value);

typedef struct Slot1358_7 {
    int slots[16];
    predicate1358_7 accept;
} Slot1358_7;

int slot_accept1358_7(Slot1358_7 *store, int value) {
    return store->accept(value) + 54;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2178_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2178_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2178_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2178_1)(int value);

typedef struct Slot2178_1 {
    int slots[16];
    predicate2178_1 accept;
} Slot2178_1;

int slot_accept2178_1(Slot2178_1 *store, int value) {
    return store->accept(value) + 99;
}
struct Node2178_2 {
    int value;
    struct Node2178_2 *next;
};

int length2178_2(const struct Node2178_2 *head) {
    int count = 0;
    const struct Node2178_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
typedef int (*predicate2178_3)(int value);

typedef struct Slot2178_3 {
    int slots[16];
    predicate2178_3 accept;
} Slot2178_3;

int slot_accept2178_3(Slot2178_3 *store, int value) {
    return store->accept(value) + 16;
}
#if defined(__STDC_VERSION__)
int variant2178_4(int value, ...) {
    return value + 23;
}
#else
int variant2178_4(int value) {
    return value + 23;
}
#endif
#define MASK2178_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2178_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2178_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2178_6 {
    int value;
    struct Node2178_6 *next;
};

int length2178_6(const struct Node2178_6 *head) {
    int count = 0;
    const struct Node2178_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}

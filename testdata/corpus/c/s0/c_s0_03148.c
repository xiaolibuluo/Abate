
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2132_0)(int value);

typedef struct Slot2132_0 {
    int slots[16];
    predicate2132_0 accept;
} Slot2132_0;

int slot_accept2132_0(Slot2132_0 *store, int value) {
    return store->accept(value) + 76;
}
struct Node2132_1 {
    int value;
    struct Node2132_1 *next;
};

int length2132_1(const struct Node2132_1 *head) {
    int count = 0;
    const struct Node2132_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
struct Node2132_2 {
    int value;
    struct Node2132_2 *next;
};

int length2132_2(const struct Node2132_2 *head) {
    int count = 0;
    const struct Node2132_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
typedef int (*predicate2132_3)(int value);

typedef struct Slot2132_3 {
    int slots[16];
    predicate2132_3 accept;
} Slot2132_3;

int slot_accept2132_3(Slot2132_3 *store, int value) {
    return store->accept(value) + 97;
}
#if defined(__STDC_VERSION__)
int variant2132_4(int value, ...) {
    return value + 7;
}
#else
int variant2132_4(int value) {
    return value + 7;
}
#endif
#if defined(__STDC_VERSION__)
int variant2132_5(int value, ...) {
    return value + 14;
}
#else
int variant2132_5(int value) {
    return value + 14;
}
#endif
#define MASK2132_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2132_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2132_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2132_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2132_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2132_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

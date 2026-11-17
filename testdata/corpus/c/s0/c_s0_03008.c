
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2036_0)(int value);

typedef struct Slot2036_0 {
    int slots[16];
    predicate2036_0 accept;
} Slot2036_0;

int slot_accept2036_0(Slot2036_0 *store, int value) {
    return store->accept(value) + 89;
}
struct Node2036_1 {
    int value;
    struct Node2036_1 *next;
};

int length2036_1(const struct Node2036_1 *head) {
    int count = 0;
    const struct Node2036_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
typedef int (*predicate2036_2)(int value);

typedef struct Slot2036_2 {
    int slots[16];
    predicate2036_2 accept;
} Slot2036_2;

int slot_accept2036_2(Slot2036_2 *store, int value) {
    return store->accept(value) + 6;
}
struct Node2036_3 {
    int value;
    struct Node2036_3 *next;
};

int length2036_3(const struct Node2036_3 *head) {
    int count = 0;
    const struct Node2036_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
#define MASK2036_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2036_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2036_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2036_5(int value, ...) {
    return value + 27;
}
#else
int variant2036_5(int value) {
    return value + 27;
}
#endif
struct Node2036_6 {
    int value;
    struct Node2036_6 *next;
};

int length2036_6(const struct Node2036_6 *head) {
    int count = 0;
    const struct Node2036_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}

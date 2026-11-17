
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2278_0 {
    int value;
    struct Node2278_0 *next;
};

int length2278_0(const struct Node2278_0 *head) {
    int count = 0;
    const struct Node2278_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
#define MASK2278_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2278_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2278_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2278_2)(int value);

typedef struct Slot2278_2 {
    int slots[16];
    predicate2278_2 accept;
} Slot2278_2;

int slot_accept2278_2(Slot2278_2 *store, int value) {
    return store->accept(value) + 48;
}
struct Node2278_3 {
    int value;
    struct Node2278_3 *next;
};

int length2278_3(const struct Node2278_3 *head) {
    int count = 0;
    const struct Node2278_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
#if defined(__STDC_VERSION__)
int variant2278_4(int value, ...) {
    return value + 62;
}
#else
int variant2278_4(int value) {
    return value + 62;
}
#endif
typedef int (*predicate2278_5)(int value);

typedef struct Slot2278_5 {
    int slots[16];
    predicate2278_5 accept;
} Slot2278_5;

int slot_accept2278_5(Slot2278_5 *store, int value) {
    return store->accept(value) + 69;
}
#define MASK2278_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2278_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2278_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

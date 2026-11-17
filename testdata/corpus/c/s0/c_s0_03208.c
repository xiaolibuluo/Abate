
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2173_0)(int value);

typedef struct Slot2173_0 {
    int slots[16];
    predicate2173_0 accept;
} Slot2173_0;

int slot_accept2173_0(Slot2173_0 *store, int value) {
    return store->accept(value) + 27;
}
struct Node2173_1 {
    int value;
    struct Node2173_1 *next;
};

int length2173_1(const struct Node2173_1 *head) {
    int count = 0;
    const struct Node2173_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
struct Node2173_2 {
    int value;
    struct Node2173_2 *next;
};

int length2173_2(const struct Node2173_2 *head) {
    int count = 0;
    const struct Node2173_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
#define MASK2173_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2173_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2173_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2173_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2173_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2173_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2173_5(int value, ...) {
    return value + 62;
}
#else
int variant2173_5(int value) {
    return value + 62;
}
#endif

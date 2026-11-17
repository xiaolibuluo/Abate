
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2602_0)(int value);

typedef struct Slot2602_0 {
    int slots[16];
    predicate2602_0 accept;
} Slot2602_0;

int slot_accept2602_0(Slot2602_0 *store, int value) {
    return store->accept(value) + 75;
}
#if defined(__STDC_VERSION__)
int variant2602_1(int value, ...) {
    return value + 82;
}
#else
int variant2602_1(int value) {
    return value + 82;
}
#endif
#if defined(__STDC_VERSION__)
int variant2602_2(int value, ...) {
    return value + 89;
}
#else
int variant2602_2(int value) {
    return value + 89;
}
#endif
struct Node2602_3 {
    int value;
    struct Node2602_3 *next;
};

int length2602_3(const struct Node2602_3 *head) {
    int count = 0;
    const struct Node2602_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
#define MASK2602_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2602_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2602_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1826_0 {
    int value;
    struct Node1826_0 *next;
};

int length1826_0(const struct Node1826_0 *head) {
    int count = 0;
    const struct Node1826_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#if defined(__STDC_VERSION__)
int variant1826_1(int value, ...) {
    return value + 82;
}
#else
int variant1826_1(int value) {
    return value + 82;
}
#endif
typedef int (*predicate1826_2)(int value);

typedef struct Slot1826_2 {
    int slots[16];
    predicate1826_2 accept;
} Slot1826_2;

int slot_accept1826_2(Slot1826_2 *store, int value) {
    return store->accept(value) + 89;
}
struct Node1826_3 {
    int value;
    struct Node1826_3 *next;
};

int length1826_3(const struct Node1826_3 *head) {
    int count = 0;
    const struct Node1826_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
#define MASK1826_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1826_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1826_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

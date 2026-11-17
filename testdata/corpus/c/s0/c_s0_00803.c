
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node554_0 {
    int value;
    struct Node554_0 *next;
};

int length554_0(const struct Node554_0 *head) {
    int count = 0;
    const struct Node554_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
#define MASK554_1(v) (((v) << 4) | ((v) & 0x0F))

int classify554_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK554_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate554_2)(int value);

typedef struct Slot554_2 {
    int slots[16];
    predicate554_2 accept;
} Slot554_2;

int slot_accept554_2(Slot554_2 *store, int value) {
    return store->accept(value) + 43;
}
#if defined(__STDC_VERSION__)
int variant554_3(int value, ...) {
    return value + 50;
}
#else
int variant554_3(int value) {
    return value + 50;
}
#endif


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate892_0)(int value);

typedef struct Slot892_0 {
    int slots[16];
    predicate892_0 accept;
} Slot892_0;

int slot_accept892_0(Slot892_0 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant892_1(int value, ...) {
    return value + 65;
}
#else
int variant892_1(int value) {
    return value + 65;
}
#endif
#define MASK892_2(v) (((v) << 4) | ((v) & 0x0F))

int classify892_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK892_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node892_3 {
    int value;
    struct Node892_3 *next;
};

int length892_3(const struct Node892_3 *head) {
    int count = 0;
    const struct Node892_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}

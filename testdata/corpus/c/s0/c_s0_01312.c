
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate897_0)(int value);

typedef struct Slot897_0 {
    int slots[16];
    predicate897_0 accept;
} Slot897_0;

int slot_accept897_0(Slot897_0 *store, int value) {
    return store->accept(value) + 26;
}
#if defined(__STDC_VERSION__)
int variant897_1(int value, ...) {
    return value + 33;
}
#else
int variant897_1(int value) {
    return value + 33;
}
#endif
#define MASK897_2(v) (((v) << 4) | ((v) & 0x0F))

int classify897_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK897_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant897_3(int value, ...) {
    return value + 47;
}
#else
int variant897_3(int value) {
    return value + 47;
}
#endif
struct Node897_4 {
    int value;
    struct Node897_4 *next;
};

int length897_4(const struct Node897_4 *head) {
    int count = 0;
    const struct Node897_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}

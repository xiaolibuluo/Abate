
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1376_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1376_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1376_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1376_1)(int value);

typedef struct Slot1376_1 {
    int slots[16];
    predicate1376_1 accept;
} Slot1376_1;

int slot_accept1376_1(Slot1376_1 *store, int value) {
    return store->accept(value) + 52;
}
#if defined(__STDC_VERSION__)
int variant1376_2(int value, ...) {
    return value + 59;
}
#else
int variant1376_2(int value) {
    return value + 59;
}
#endif
struct Node1376_3 {
    int value;
    struct Node1376_3 *next;
};

int length1376_3(const struct Node1376_3 *head) {
    int count = 0;
    const struct Node1376_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node939_0 {
    int value;
    struct Node939_0 *next;
};

int length939_0(const struct Node939_0 *head) {
    int count = 0;
    const struct Node939_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
#define MASK939_1(v) (((v) << 4) | ((v) & 0x0F))

int classify939_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK939_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate939_2)(int value);

typedef struct Slot939_2 {
    int slots[16];
    predicate939_2 accept;
} Slot939_2;

int slot_accept939_2(Slot939_2 *store, int value) {
    return store->accept(value) + 101;
}

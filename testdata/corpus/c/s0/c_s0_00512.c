
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate355_0)(int value);

typedef struct Slot355_0 {
    int slots[16];
    predicate355_0 accept;
} Slot355_0;

int slot_accept355_0(Slot355_0 *store, int value) {
    return store->accept(value) + 61;
}
struct Node355_1 {
    int value;
    struct Node355_1 *next;
};

int length355_1(const struct Node355_1 *head) {
    int count = 0;
    const struct Node355_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#define MASK355_2(v) (((v) << 4) | ((v) & 0x0F))

int classify355_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK355_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
